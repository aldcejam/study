package commands

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"study_manager/src/infra/database"

	"github.com/ostafen/clover/v2/query"
)

// HandleVerNota busca e envia o conteúdo de uma nota específica
func HandleVerNota(dbPath string, shortID string, sendFunc func(int64, string, string) error, chatID int64) {
	db, err := database.InitDB(dbPath)
	if err != nil {
		return
	}
	defer db.Close()

	doc, err := db.FindFirst(query.NewQuery("notes").Where(query.Field("short_id").Eq(shortID)))
	if err != nil || doc == nil {
		sendFunc(chatID, "❌ Nota não encontrada.", "")
		return
	}

	var note database.NoteDoc
	doc.Unmarshal(&note)

	// Localizar o arquivo físico
	vaultRoot, _ := filepath.Abs("..")
	fullPath := filepath.Join(vaultRoot, note.RelativePath, note.Filename)

	content, err := os.ReadFile(fullPath)
	if err != nil {
		log.Printf("Read error: %v", err)
		sendFunc(chatID, fmt.Sprintf("❌ Erro ao ler arquivo: %s", note.Filename), "")
		return
	}

	// Enviamos APENAS o conteúdo original
	if err := sendFunc(chatID, string(content), ""); err != nil {
		log.Printf("Send content error: %v", err)
	}
}
