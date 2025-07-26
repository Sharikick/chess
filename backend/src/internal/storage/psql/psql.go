// Package psql ...
package psql

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	pool *pgxpool.Pool
}

func New(url string) (*Storage, error) {
	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		return nil, err
	}
	defer pool.Close()
	return &Storage{pool:pool}, nil
}






