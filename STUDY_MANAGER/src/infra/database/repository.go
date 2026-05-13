package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

func (r *Repository) Close() {
	r.pool.Close()
}

// Modelos simplificados para as queries

type Note struct {
	ID           int
	Filename     string
	RelativePath string
	ShortID      string
	Tema         string
	Subtema      *string
	Revisoes     []byte
	References   []byte
	Homework     []byte
	Tags         []string
	Activity     *string
	UpdatedAt    *time.Time
}

// GetNoteByShortID retorna os principais campos de uma nota
func (r *Repository) GetNoteByShortID(ctx context.Context, shortID string) (*Note, error) {
	var note Note
	err := r.pool.QueryRow(ctx, "SELECT relative_path, filename, tema, homework, \"references\" FROM notes WHERE short_id = $1", shortID).
		Scan(&note.RelativePath, &note.Filename, &note.Tema, &note.Homework, &note.References)
	if err != nil {
		return nil, err
	}
	note.ShortID = shortID
	return &note, nil
}

// GetAllNotesInfo retorna informações básicas de todas as notas
func (r *Repository) GetAllNotesInfo(ctx context.Context) ([]Note, error) {
	rows, err := r.pool.Query(ctx, "SELECT relative_path, tema, short_id FROM notes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []Note
	for rows.Next() {
		var n Note
		if err := rows.Scan(&n.RelativePath, &n.Tema, &n.ShortID); err != nil {
			continue
		}
		notes = append(notes, n)
	}
	return notes, nil
}

// UpsertStudySession cria a sessão de estudo se não existir
func (r *Repository) UpsertStudySession(ctx context.Context, chatID int64, threadID int, shortID string) error {
	_, err := r.pool.Exec(ctx, `
		INSERT INTO study_sessions (chat_id, thread_id, short_id, history)
		VALUES ($1, $2, $3, '[]'::jsonb)
		ON CONFLICT (chat_id, short_id) DO UPDATE SET
			thread_id = EXCLUDED.thread_id,
			history = EXCLUDED.history
	`, chatID, threadID, shortID)
	return err
}

// GetStudySession retorna os dados de uma sessão de estudo ativa. Se for um fork, resolve o shortID buscando recursivamente o pai.
func (r *Repository) GetStudySession(ctx context.Context, chatID int64, threadID int) (string, []byte, error) {
	var shortID *string
	var parentThreadID *int
	var history []byte
	
	err := r.pool.QueryRow(ctx, "SELECT short_id, parent_thread_id, history FROM study_sessions WHERE chat_id = $1 AND thread_id = $2", chatID, threadID).
		Scan(&shortID, &parentThreadID, &history)
	
	if err != nil {
		return "", nil, err
	}

	finalShortID := ""
	if shortID != nil {
		finalShortID = *shortID
	} else if parentThreadID != nil {
		finalShortID, _, err = r.GetStudySession(ctx, chatID, *parentThreadID)
		if err != nil {
			return "", nil, fmt.Errorf("erro ao buscar sessão pai: %w", err)
		}
	} else {
		return "", nil, fmt.Errorf("sessão corrompida: sem short_id e sem pai")
	}

	return finalShortID, history, nil
}

// CreateForkSession clona uma sessão de estudos, apontando para o pai e preservando o histórico.
func (r *Repository) CreateForkSession(ctx context.Context, chatID int64, parentThreadID int, newThreadID int, history []byte) error {
	_, err := r.pool.Exec(ctx, `
		INSERT INTO study_sessions (chat_id, thread_id, parent_thread_id, history)
		VALUES ($1, $2, $3, $4)
	`, chatID, newThreadID, parentThreadID, history)
	return err
}

// GetThreadIDForNote retorna o thread_id de um tópico existente para a nota no chat
func (r *Repository) GetThreadIDForNote(ctx context.Context, chatID int64, shortID string) (int, error) {
	var threadID int
	err := r.pool.QueryRow(ctx, "SELECT thread_id FROM study_sessions WHERE chat_id = $1 AND short_id = $2 LIMIT 1", chatID, shortID).
		Scan(&threadID)
	return threadID, err
}

// DeleteStudySession deleta a sessão de estudos associada a um tópico que não existe mais
func (r *Repository) DeleteStudySession(ctx context.Context, chatID int64, threadID int) error {
	_, err := r.pool.Exec(ctx, "DELETE FROM study_sessions WHERE chat_id = $1 AND thread_id = $2", chatID, threadID)
	return err
}

// UpdateStudySessionHistory atualiza o JSON de histórico do gemini
func (r *Repository) UpdateStudySessionHistory(ctx context.Context, chatID int64, threadID int, history []byte) error {
	_, err := r.pool.Exec(ctx, "UPDATE study_sessions SET history = $1 WHERE chat_id = $2 AND thread_id = $3", history, chatID, threadID)
	return err
}

// UpsertNote insere ou atualiza uma nota completa
func (r *Repository) UpsertNote(ctx context.Context, n Note) error {
	_, err := r.pool.Exec(ctx, `
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
	`, n.Filename, n.RelativePath, n.ShortID, n.Tema, n.Subtema, n.Revisoes, n.References, n.Homework, n.Tags, n.Activity, n.UpdatedAt)
	return err
}

// GetNotificationStatus verifica o status da notificação
func (r *Repository) GetNotificationStatus(ctx context.Context, relativePath string) (*time.Time, bool, error) {
	var lastNotifiedAt *time.Time
	var completed bool
	err := r.pool.QueryRow(ctx, "SELECT last_notified_at, completed FROM notifications WHERE relative_path = $1", relativePath).
		Scan(&lastNotifiedAt, &completed)
	if err == pgx.ErrNoRows {
		return nil, false, nil // não encontrou, consideramos null/false
	}
	return lastNotifiedAt, completed, err
}

// UpsertNotification salva o log da notificação enviada
func (r *Repository) UpsertNotification(ctx context.Context, relativePath string, lastNotifiedAt *time.Time, completed bool) error {
	_, err := r.pool.Exec(ctx, `
		INSERT INTO notifications (relative_path, last_notified_at, completed)
		VALUES ($1, $2, $3)
		ON CONFLICT (relative_path) DO UPDATE SET
			last_notified_at = EXCLUDED.last_notified_at,
			completed = EXCLUDED.completed
	`, relativePath, lastNotifiedAt, completed)
	return err
}
