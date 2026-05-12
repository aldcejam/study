package commands

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"study_manager/src/infra/database"
)

// HandleVerNota busca e envia o conteúdo de uma nota específica
func HandleVerNota(repo *database.Repository, shortID string, sendFunc func(int64, string, string) error, chatID int64) {
	note, err := repo.GetNoteByShortID(context.Background(), shortID)
	if err != nil {
		log.Printf("Erro ao buscar nota %s: %v", shortID, err)
		sendFunc(chatID, "❌ Nota não encontrada.", "")
		return
	}

	// Localizar o arquivo físico usando a variável de ambiente do Docker ou fallback local
	vaultRoot := os.Getenv("VAULT_PATH")
	if vaultRoot == "" {
		vaultRoot, _ = filepath.Abs("..")
	}
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
