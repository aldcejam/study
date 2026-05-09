package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"study_manager/src/infra/database"
	"study_manager/src/utils/dates"
	"study_manager/src/utils/models"

	"github.com/ostafen/clover/v2/document"
	"github.com/ostafen/clover/v2/query"
)

const bigFuture = 9999

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func shouldNotify(note models.SummaryNotificationOutput, notifDoc *database.NotificationDoc) bool {
	if notifDoc == nil || notifDoc.LastNotifiedAt == nil {
		return true
	}

	lastNotified, err1 := time.Parse(time.RFC3339, *notifDoc.LastNotifiedAt)
	updatedAt, err2 := time.Parse(time.RFC3339, note.UpdatedAt)

	if err1 != nil || err2 != nil {
		return true
	}

	// Compara via Unix para evitar problemas de precisão/fuso no parsing
	if updatedAt.Unix() > lastNotified.Unix() {
		return true
	}

	// Regra de 48h
	if time.Since(lastNotified).Hours() >= 48 {
		return true
	}

	return false
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

	selfDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	// Sobe um nível para sair de bin/ e chegar na raiz do projeto, depois entra em output/
	dbPath := filepath.Join(selfDir, "..", "output", "clover_db")
	db, err := database.InitDB(dbPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "summaryNotification: erro ao inicializar CloverDB:", err)
		os.Exit(1)
	}
	defer db.Close()

	var urgentes []models.SummaryNotificationOutput
	var aNotificar []models.SummaryNotificationOutput
	today := time.Now().Truncate(24 * time.Hour)
	nowStr := time.Now().Format(time.RFC3339)

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

		// 3. Atualizar tabela `notes`
		noteDocs, _ := db.FindAll(query.NewQuery("notes").Where(query.Field("relative_path").Eq(note.RelativePath)))
		var createdAt time.Time
		if len(noteDocs) > 0 {
			var existingNote database.NoteDoc
			noteDocs[0].Unmarshal(&existingNote)
			createdAt = existingNote.CreatedAt
		} else {
			createdAt = time.Now()
		}

		newNoteDoc := database.NoteDoc{
			ScannerOutput: note,
			ShortID:       shortID,
			CreatedAt:     createdAt,
		}

		docMap := make(map[string]interface{})
		b, _ := json.Marshal(newNoteDoc)
		json.Unmarshal(b, &docMap)
		docMap["short_id"] = shortID // Garantia extra

		if len(noteDocs) > 0 {
			err := db.Update(query.NewQuery("notes").Where(query.Field("_id").Eq(noteDocs[0].ObjectId())), docMap)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERRO ao atualizar nota: %v\n", err)
			}
		} else {
			doc := document.NewDocumentOf(docMap)
			_, err := db.InsertOne("notes", doc)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERRO ao inserir nota: %v\n", err)
			}
		}

		// 4. Checar tabela `notifications`
		var notifDoc *database.NotificationDoc
		notifDocs, _ := db.FindAll(query.NewQuery("notifications").Where(query.Field("relative_path").Eq(note.RelativePath)))
		if len(notifDocs) > 0 {
			var existingNotif database.NotificationDoc
			notifDocs[0].Unmarshal(&existingNotif)
			notifDoc = &existingNotif
		}

		shouldNotifyNow := false
		if !completed {
			if status == "ATRASADA" || status == "HOJE" {
				urgentes = append(urgentes, pn)
				shouldNotifyNow = shouldNotify(pn, notifDoc)
			}
		}

		// 5. Preparar para notificar e atualizar `notifications`
		if shouldNotifyNow {
			aNotificar = append(aNotificar, pn)
		}

		newNotif := database.NotificationDoc{
			RelativePath:   note.RelativePath,
			Completed:      completed,
			LastNotifiedAt: nil,
		}

		if notifDoc != nil {
			newNotif.LastNotifiedAt = notifDoc.LastNotifiedAt
		}

		if shouldNotifyNow {
			// Update the time since we will notify
			newNotif.LastNotifiedAt = &nowStr
		}

		notifMap := make(map[string]interface{})
		b2, _ := json.Marshal(newNotif)
		json.Unmarshal(b2, &notifMap)

		if len(notifDocs) > 0 {
			err := db.Update(query.NewQuery("notifications").Where(query.Field("_id").Eq(notifDocs[0].ObjectId())), notifMap)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERRO ao atualizar notificação: %v\n", err)
			}
		} else {
			doc := document.NewDocumentOf(notifMap)
			_, err := db.InsertOne("notifications", doc)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERRO ao inserir notificação: %v\n", err)
			}
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
