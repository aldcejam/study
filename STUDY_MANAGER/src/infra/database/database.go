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

	CREATE TABLE IF NOT EXISTS study_sessions (
		chat_id BIGINT NOT NULL,
		thread_id INTEGER NOT NULL,
		short_id TEXT NOT NULL REFERENCES notes(short_id) ON DELETE CASCADE,
		history JSONB DEFAULT '[]'::jsonb, 
		PRIMARY KEY (chat_id, thread_id) 
	);
	`

	_, err := pool.Exec(ctx, schema)
	return err
}
