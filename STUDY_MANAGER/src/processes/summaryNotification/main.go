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

	if updatedAt.After(lastNotified) {
		return true
	}

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
	sb.WriteString("📚 *RESUMO DE ESTUDOS*\n")
	sb.WriteString("──────────────────\n\n")

	for _, note := range aNotificar {
		icon := "📅"
		atrasoStr := "HOJE"
		if note.StatusRevisao == "ATRASADA" {
			icon = "🚨"
			atrasoStr = fmt.Sprintf("%dd atrás", abs(note.DiasAtraso))
		}

		parts := strings.Split(note.RelativePath, string(filepath.Separator))
		contexto := parts[len(parts)-1]

		sb.WriteString(fmt.Sprintf("%s *%s*\n", icon, note.Tema))
		sb.WriteString(fmt.Sprintf("   └ 📂 %s | ⏳ %s\n", contexto, atrasoStr))

		if note.Activity != nil && *note.Activity != "" {
			sb.WriteString(fmt.Sprintf("   📝 *Atividade:* %s\n", *note.Activity))
		}

		if note.References != nil {
			switch refs := note.References.(type) {
			case []interface{}:
				for _, ref := range refs {
					switch r := ref.(type) {
					case map[string]interface{}:
						desc, _ := r["description"].(string)
						src, _ := r["source"].(string)
						sb.WriteString(fmt.Sprintf("   🔗 _%s: %s_\n", desc, src))
					case string:
						sb.WriteString(fmt.Sprintf("   🔗 _%s_\n", r))
					}
				}
			}
			sb.WriteString("\n")
		} else {
			sb.WriteString("\n")
		}
	}

	sb.WriteString("──────────────────\n")
	sb.WriteString("_Abra seu Obsidian para revisar!_")

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
		// 1. Calcular status e diasAtraso
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
			ID:            note.RelativePath, // Using relative_path as ID for tracking
			DiasAtraso:    diasAtraso,
			StatusRevisao: status,
		}

		// 2. Atualizar tabela `notes`
		noteDocs, _ := db.FindAll(query.NewQuery("notes").Where(query.Field("relative_path").Eq(note.RelativePath)))
		var createdAt time.Time
		if len(noteDocs) > 0 {
			var existingNote database.NoteDoc
			noteDocs[0].Unmarshal(&existingNote)
			createdAt = existingNote.CreatedAt
		} else {
			createdAt = time.Now()
		}

		hash := md5.Sum([]byte(note.RelativePath))
		shortID := hex.EncodeToString(hash[:3]) // 6 caracteres

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
			doc := document.NewDocumentOf(docMap)
			db.ReplaceById("notes", noteDocs[0].ObjectId(), doc)
		} else {
			doc := document.NewDocumentOf(docMap)
			db.InsertOne("notes", doc)
		}

		// 3. Checar tabela `notifications`
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

		// 4. Preparar para notificar e atualizar `notifications`
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
			doc := document.NewDocumentOf(notifMap)
			db.ReplaceById("notifications", notifDocs[0].ObjectId(), doc)
		} else {
			doc := document.NewDocumentOf(notifMap)
			db.InsertOne("notifications", doc)
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
