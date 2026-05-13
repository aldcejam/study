package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// InitDB inicializa o banco PostgreSQL usando o driver pgx/v5 e um Pool de conexões.
// Executa o schema (migrations automáticas) se necessário.
func InitDB(connString string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer parse do DATABASE_URL: %w", err)
	}

	// Configurações do Pool
	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnIdleTime = 5 * time.Minute

	ctx := context.Background()
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("banco de dados não está respondendo: %w", err)
	}

	if err := RunMigrations(ctx, pool); err != nil {
		return nil, fmt.Errorf("erro ao executar migrations: %w", err)
	}

	return pool, nil
}

// RunMigrations executa o SQL de criação das tabelas
func RunMigrations(ctx context.Context, pool *pgxpool.Pool) error {
	schema := `
	CREATE TABLE IF NOT EXISTS notes (
		id SERIAL PRIMARY KEY,
		filename TEXT NOT NULL,
		relative_path TEXT UNIQUE NOT NULL,
		short_id TEXT UNIQUE NOT NULL,
		tema TEXT NOT NULL,
		subtema TEXT,
		revisoes JSONB DEFAULT '[]'::jsonb, 
		"references" JSONB DEFAULT '[]'::jsonb,
		homework JSONB DEFAULT '[]'::jsonb,
		tags TEXT[] DEFAULT '{}',
		activity TEXT,
		updated_at TIMESTAMP WITH TIME ZONE,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS notifications (
		relative_path TEXT PRIMARY KEY REFERENCES notes(relative_path) ON DELETE CASCADE,
		last_notified_at TIMESTAMP WITH TIME ZONE,
		completed BOOLEAN DEFAULT FALSE
	);

	DO $$ 
	BEGIN 
		-- Create the table first to ensure it exists
		CREATE TABLE IF NOT EXISTS study_sessions (
			chat_id BIGINT NOT NULL,
			thread_id INTEGER NOT NULL,
			short_id TEXT REFERENCES notes(short_id) ON DELETE CASCADE,
			parent_thread_id INTEGER,
			history JSONB DEFAULT '[]'::jsonb, 
			PRIMARY KEY (chat_id, thread_id),
			CONSTRAINT unique_chat_note UNIQUE (chat_id, short_id),
			CONSTRAINT fk_parent FOREIGN KEY (chat_id, parent_thread_id) REFERENCES study_sessions(chat_id, thread_id) ON DELETE CASCADE
		);
		
		-- If the table was already there, apply alterations:
		ALTER TABLE study_sessions ALTER COLUMN short_id DROP NOT NULL;
		
		IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='study_sessions' AND column_name='parent_thread_id') THEN
			ALTER TABLE study_sessions ADD COLUMN parent_thread_id INTEGER;
			ALTER TABLE study_sessions ADD CONSTRAINT fk_parent FOREIGN KEY (chat_id, parent_thread_id) REFERENCES study_sessions(chat_id, thread_id) ON DELETE CASCADE;
		END IF;
	END $$;
	`

	_, err := pool.Exec(ctx, schema)
	return err
}
