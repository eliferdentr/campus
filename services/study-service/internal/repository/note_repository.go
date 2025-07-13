package repository

import (
	"context"

	"study-service/internal/domain" // Projenizin go.mod dosyasındaki module adının 'study-service' olduğu varsayılmıştır.

	"github.com/jackc/pgx/v5/pgxpool"
)

// NoteRepository, not verileri üzerindeki operasyonlar için arayüzü tanımlar.
type NoteRepository interface {
	// Save, verilen notu veritabanına kaydeder.
	Save(ctx context.Context, note *domain.Note) error
}

// postgresNoteRepository, NoteRepository arayüzünün PostgreSQL implementasyonudur.
type postgresNoteRepository struct {
	db *pgxpool.Pool
}

// NewPostgresNoteRepository, postgresNoteRepository'nin yeni bir örneğini oluşturur.
func NewPostgresNoteRepository(db *pgxpool.Pool) NoteRepository {
	return &postgresNoteRepository{db: db}
}

// Save, yeni bir notu veritabanına kaydeder.
func (r *postgresNoteRepository) Save(ctx context.Context, note *domain.Note) error {
	// TODO: SQL INSERT sorgusu buraya gelecek
	return nil
}
