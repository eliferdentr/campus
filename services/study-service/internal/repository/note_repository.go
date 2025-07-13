package repository

import (
	"context"

	"campus-project.com/study-service/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NoteRepository, not verileri işlemleri için arayüzü tanımlar.
type NoteRepository interface {
	Save(ctx context.Context, note *domain.Note) error
}

type postgresNoteRepository struct {
	db *pgxpool.Pool
}

// NewPostgresNoteRepository, postgresNoteRepository için yeni bir örnek oluşturur.
func NewPostgresNoteRepository(db *pgxpool.Pool) NoteRepository {
	return &postgresNoteRepository{db: db}
}

func (r *postgresNoteRepository) Save(ctx context.Context, note *domain.Note) error {
	// TODO: SQL INSERT sorgusu buraya gelecek
	return nil
}
