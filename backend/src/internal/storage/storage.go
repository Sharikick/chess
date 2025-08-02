package storage

import (
	"chess/internal/config"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(config config.DatabaseConfig) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), config.GetConn())
	if err != nil {
		return nil, err
	}
	return pool, nil
}
