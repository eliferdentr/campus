package repository

import (
	"context"

	"study-service/internal/domain" // Projenizin go.mod dosyasındaki module adının 'study-service' olduğu varsayılmıştır.
	"study-service/internal/utils"

	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NoteRepository, not verileri üzerindeki operasyonlar için arayüzü tanımlar.
type NoteRepository interface {
	// Save, verilen notu veritabanına kaydeder.
	Create(ctx context.Context, note *domain.Note) (error, string)
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
func (r *postgresNoteRepository) Create(ctx context.Context, note *domain.Note) (error, string) {
	uuid := utils.GenerateUUID()
	if uuid == "" {
		return fmt.Errorf("unable to generate uuid"), ""
	}
	if note.ID == "" {
		note.ID = uuid
	}
	row := r.db.QueryRow(
		`INSERT INTO notes (id, title, description, file_path, user_id, university_id, course_code) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		note.ID, note.Title, note.Description, note.FilePath, note.UserID, note.UniversityID, note.CourseCode)
	err := row.Scan(&note.ID)

	return nil, uuid
}
