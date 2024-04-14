package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func NewClient(postgresUrl string) (*pgxpool.Pool, error) {
	ctx := context.Background()
	db, err := pgxpool.New(ctx, postgresUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	return db, nil
}
