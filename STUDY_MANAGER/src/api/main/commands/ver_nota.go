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
func HandleVerNota(connStr string, shortID string, sendFunc func(int64, string, string) error, chatID int64) {
	pool, err := database.InitDB(connStr)
	if err != nil {
		log.Printf("DB error: %v", err)
		return
	}
	defer pool.Close()

	var relativePath, filename string
	err = pool.QueryRow(context.Background(), "SELECT relative_path, filename FROM notes WHERE short_id = $1", shortID).Scan(&relativePath, &filename)
	if err != nil {
		sendFunc(chatID, "❌ Nota não encontrada.", "")
		return
	}

	// Localizar o arquivo físico usando a variável de ambiente do Docker ou fallback local
	vaultRoot := os.Getenv("VAULT_PATH")
	if vaultRoot == "" {
		vaultRoot, _ = filepath.Abs("..")
	}
	fullPath := filepath.Join(vaultRoot, relativePath, filename)

	content, err := os.ReadFile(fullPath)
	if err != nil {
		log.Printf("Read error: %v", err)
		sendFunc(chatID, fmt.Sprintf("❌ Erro ao ler arquivo: %s", filename), "")
		return
	}

	// Enviamos APENAS o conteúdo original
	if err := sendFunc(chatID, string(content), ""); err != nil {
		log.Printf("Send content error: %v", err)
	}
}
