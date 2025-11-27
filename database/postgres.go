package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Postgres *pgxpool.Pool

func InitPostgres() error {
	dsn := os.Getenv("POSTGRES_URL")
	if dsn == "" {
		return fmt.Errorf("POSTGRES_URL is empty")
	}

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return fmt.Errorf("failed to connect PostgreSQL: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return fmt.Errorf("failed to ping PostgreSQL: %w", err)
	}

	Postgres = pool
	fmt.Println("Connected to PostgreSQL")
	return nil
}
