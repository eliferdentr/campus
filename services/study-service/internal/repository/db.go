package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

// NewConnection, yeni bir veritabanı bağlantısı havuzu oluşturur.
func NewConnection(dbURL string) (pool *pgxpool.Pool, err error) {
	pool, err = pgxpool.New(context.Background(), dbURL)
	if err != nil {
		return nil, err
	}
	//bağlantıyı test et
	err = pool.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return pool, nil
}
