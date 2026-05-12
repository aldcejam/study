package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"

	"study_manager/src/infra/database"
	"study_manager/src/utils/dates"
	"study_manager/src/utils/models"
)

const bigFuture = 9999

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}



func formatMessage(aNotificar []models.SummaryNotificationOutput) string {
	if len(aNotificar) == 0 {
		return "SEM_ALTERACOES"
	}

	sort.Slice(aNotificar, func(i, j int) bool {
		return aNotificar[i].DiasAtraso < aNotificar[j].DiasAtraso
	})

	var sb strings.Builder
	sb.WriteString("<b>📚 RESUMO DE ESTUDOS</b>\n")
	sb.WriteString("──────────────────\n\n")

	today := time.Now().Truncate(24 * time.Hour)

	for _, note := range aNotificar {
		// Escape simples para evitar quebra de HTML
		safeTema := strings.ReplaceAll(note.Tema, "<", "&lt;")
		safeTema = strings.ReplaceAll(safeTema, ">", "&gt;")

		sb.WriteString(fmt.Sprintf("🚨 <b>%s</b>\n", safeTema))
		sb.WriteString(fmt.Sprintf("   └ 📂 %s\n", note.RelativePath))
		sb.WriteString("   └ 📅 Revisões: \n")

		// Listar todas as revisões e seus status
		for _, rev := range note.Revisoes {
			statusStr := ""
			if strings.TrimSpace(rev.Status) == "x" {
				statusStr = "Feito"
			} else {
				d := dates.ParseDate(rev.Data)
				diff := int(today.Sub(d.Truncate(24 * time.Hour)).Hours() / 24)
				
				if diff > 0 {
					statusStr = fmt.Sprintf("Atrasado %d dias", diff)
				} else if diff == 0 {
					statusStr = "Hoje"
				} else {
					statusStr = "Pendente"
				}
			}
			sb.WriteString(fmt.Sprintf("           %s: %s\n", rev.Data, statusStr))
		}

		sb.WriteString(fmt.Sprintf("   └ 📱 Menu: /ver_%s\n\n", note.ShortID))
	}

	sb.WriteString("──────────────────\n")
	sb.WriteString("<i>Abra seu menu para revisar!</i>")

	return sb.String()
}

func main() {
	raw, err := io.ReadAll(os.Stdin)
	if err != nil || len(strings.TrimSpace(string(raw))) == 0 {
		fmt.Fprintln(os.Stderr, "summaryNotification: nenhum dado recebido via stdin")
		os.Exit(1)
	}

	var notes []models.ScannerOutput
	if err := json.Unmarshal(raw, &notes); err != nil {
		fmt.Fprintln(os.Stderr, "summaryNotification: erro ao parsear JSON:", err)
		os.Exit(1)
	}

	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgres://study_user:study_password@localhost:5432/study_db?sslmode=disable"
	}
	pool, err := database.InitDB(connStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, "summaryNotification: erro ao inicializar PostgreSQL:", err)
		os.Exit(1)
	}
	defer pool.Close()

	var urgentes []models.SummaryNotificationOutput
	var aNotificar []models.SummaryNotificationOutput
	today := time.Now().Truncate(24 * time.Hour)

	for _, note := range notes {
		// 1. Calcular ShortID (usado no banco e na notificação)
		hash := md5.Sum([]byte(note.RelativePath))
		shortID := hex.EncodeToString(hash[:3]) // 6 caracteres
		note.ShortID = shortID

		// 2. Calcular status e diasAtraso
		diasAtraso := bigFuture
		status := "FUTURA"
		var pendentes []models.Revision

		for _, r := range note.Revisoes {
			if r.Status == " " {
				pendentes = append(pendentes, r)
			}
		}

		completed := len(pendentes) == 0

		if completed {
			status = "EM_DIA"
			diasAtraso = 0
		} else {
			for _, rev := range pendentes {
				d := dates.ParseDate(rev.Data)
				diff := int(d.Truncate(24*time.Hour).Sub(today).Hours() / 24)
				if diff < diasAtraso {
					diasAtraso = diff
				}
			}
			switch {
			case diasAtraso < 0:
				status = "ATRASADA"
			case diasAtraso == 0:
				status = "HOJE"
			default:
				status = "FUTURA"
			}
		}

		pn := models.SummaryNotificationOutput{
			ScannerOutput: note,
			ID:            note.RelativePath,
			DiasAtraso:    diasAtraso,
			StatusRevisao: status,
		}

		// 3. Atualizar tabela `notes` (Upsert)
		var updatedAtTime *time.Time
		if t, err := time.Parse(time.RFC3339, note.UpdatedAt); err == nil {
			updatedAtTime = &t
		}
		
		revisoesJson, _ := json.Marshal(note.Revisoes)
		referencesJson, _ := json.Marshal(note.References)
		homeworkJson, _ := json.Marshal(note.Homework)

		_, err := pool.Exec(context.Background(), `
			INSERT INTO notes (filename, relative_path, short_id, tema, subtema, revisoes, "references", homework, tags, activity, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
			ON CONFLICT (relative_path) DO UPDATE SET
				filename = EXCLUDED.filename,
				short_id = EXCLUDED.short_id,
				tema = EXCLUDED.tema,
				subtema = EXCLUDED.subtema,
				revisoes = EXCLUDED.revisoes,
				"references" = EXCLUDED."references",
				homework = EXCLUDED.homework,
				tags = EXCLUDED.tags,
				activity = EXCLUDED.activity,
				updated_at = EXCLUDED.updated_at
		`, note.Filename, note.RelativePath, shortID, note.Tema, note.Subtema, revisoesJson, referencesJson, homeworkJson, note.Tags, note.Activity, updatedAtTime)

		if err != nil {
			fmt.Fprintf(os.Stderr, "ERRO ao atualizar nota %s: %v\n", note.RelativePath, err)
		}

		// 4. Checar tabela `notifications`
		var lastNotifiedAt *time.Time
		var notifCompleted bool

		err = pool.QueryRow(context.Background(), "SELECT last_notified_at, completed FROM notifications WHERE relative_path = $1", note.RelativePath).Scan(&lastNotifiedAt, &notifCompleted)

		shouldNotifyNow := false
		if !completed {
			if status == "ATRASADA" || status == "HOJE" {
				urgentes = append(urgentes, pn)
				
				if lastNotifiedAt == nil {
					shouldNotifyNow = true
				} else {
					if updatedAtTime != nil && updatedAtTime.Unix() > lastNotifiedAt.Unix() {
						shouldNotifyNow = true
					} else if time.Since(*lastNotifiedAt).Hours() >= 48 {
						shouldNotifyNow = true
					}
				}
			}
		}

		// 5. Preparar para notificar e atualizar `notifications`
		if shouldNotifyNow {
			aNotificar = append(aNotificar, pn)
			now := time.Now()
			lastNotifiedAt = &now
		}

		// UPSERT notification
		_, err = pool.Exec(context.Background(), `
			INSERT INTO notifications (relative_path, last_notified_at, completed)
			VALUES ($1, $2, $3)
			ON CONFLICT (relative_path) DO UPDATE SET
				last_notified_at = EXCLUDED.last_notified_at,
				completed = EXCLUDED.completed
		`, note.RelativePath, lastNotifiedAt, completed)

		if err != nil {
			fmt.Fprintf(os.Stderr, "ERRO ao atualizar notificação %s: %v\n", note.RelativePath, err)
		}
	}

	var msg string
	if len(urgentes) == 0 {
		msg = "SEM_REVISOES"
	} else {
		msg = formatMessage(aNotificar)
	}

	fmt.Print(msg)
}
