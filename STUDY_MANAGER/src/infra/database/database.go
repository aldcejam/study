package database

import (
	"os"

	"github.com/ostafen/clover/v2"
)

// InitDB inicializa o banco CloverDB e garante que as tabelas necessárias existam.
func InitDB(dbPath string) (*clover.DB, error) {
	if err := os.MkdirAll(dbPath, 0755); err != nil {
		return nil, err
	}
	db, err := clover.Open(dbPath)
	if err != nil {
		return nil, err
	}

	// Criação de coleções caso não existam
	hasNotes, _ := db.HasCollection("notes")
	if !hasNotes {
		db.CreateCollection("notes")
	}
	hasNotif, _ := db.HasCollection("notifications")
	if !hasNotif {
		db.CreateCollection("notifications")
	}
	hasStudySessions, _ := db.HasCollection("study_sessions")
	if !hasStudySessions {
		db.CreateCollection("study_sessions")
	}

	// Garantir que os índices existem para melhor performance e buscas
	db.CreateIndex("notes", "relative_path")
	db.CreateIndex("notes", "updatedAt")
	db.CreateIndex("notifications", "relative_path")

	return db, nil
}
