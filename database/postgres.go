package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Postgres *pgxpool.Pool  // INI YANG DIPANGGIL DI main.go

func InitPostgres() error {
	dsn := os.Getenv("POSTGRES_URL")
	if dsn == "" {
		return fmt.Errorf("POSTGRES_URL is empty")
	}

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return err
	}

	Postgres = pool
	fmt.Println("Connected to PostgreSQL")
	return nil
}
