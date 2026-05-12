package database

import (
	"time"

	"study_manager/src/utils/models"
)

// NoteDoc define o schema utilizado para armazenar os metadados de uma nota no Postgres.
type NoteDoc struct {
	models.ScannerOutput
	ShortID   string    `json:"short_id" db:"short_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// NotificationDoc define o schema para armazenar o histórico e controle de notificações no Postgres.
type NotificationDoc struct {
	RelativePath   string     `json:"relative_path" db:"relative_path"`
	LastNotifiedAt *time.Time `json:"last_notified_at" db:"last_notified_at"`
	Completed      bool       `json:"completed" db:"completed"`
}

// StudySessionDoc representa a sessão do Gemini CLI
type StudySessionDoc struct {
	ChatID   int64    `json:"chat_id" db:"chat_id"`
	ThreadID int      `json:"thread_id" db:"thread_id"`
	ShortID  string   `json:"short_id" db:"short_id"`
	History  []string `json:"history" db:"history"`
}
