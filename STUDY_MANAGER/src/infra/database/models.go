package database

import (
	"time"

	"study_manager/src/utils/models"
)

// NoteDoc define o schema utilizado para armazenar os metadados de uma nota no banco CloverDB.
type NoteDoc struct {
	models.ScannerOutput
	ShortID   string    `json:"short_id"`
	CreatedAt time.Time `json:"createdAt"`
}

// NotificationDoc define o schema para armazenar o histórico e controle de notificações no banco.
type NotificationDoc struct {
	RelativePath   string  `json:"relative_path"`
	LastNotifiedAt *string `json:"last_notified_at"`
	Completed      bool    `json:"completed"`
}
