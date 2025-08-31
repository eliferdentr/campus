package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"study-service/internal/domain" // Projenizin go.mod dosyasındaki module adının 'study-service' olduğu varsayılmıştır.

	"github.com/jackc/pgx/v5"
)

var ErrNoteNotFound = errors.New("not bulunamadı")

// postgresNoteRepository, NoteRepository arayüzünün PostgreSQL implementasyonudur.
type postgresNoteRepository struct {
	db DBPool
}

// NewPostgresNoteRepository, postgresNoteRepository'nin yeni bir örneğini oluşturur.
func NewPostgresNoteRepository(db DBPool) NoteRepository {
	return &postgresNoteRepository{db: db}
}

// sadece save işlemi ile ilgilenir, verinin doğrulupunu kontrol etmez
func (r *postgresNoteRepository) Create(ctx context.Context, note *domain.Note) (string, error) {
	var uuid string
	row := r.db.QueryRow(
		ctx,
		`INSERT INTO notes (title, description, file_path, user_id, university_id, course_code) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		note.Title, note.Description, note.FilePath, note.UserID, note.UniversityID, note.CourseCode)
	err := row.Scan(&uuid)
	if err != nil {
		return "", fmt.Errorf("not oluşturulurken hata oluştu: %w", err)
	}
	return uuid, nil
}

// sadece update işlemi ile ilgilenir, verinin doğruluğu ile ilgilenmez
func (r *postgresNoteRepository) Update(ctx context.Context, note *domain.Note) (string, error) {
	query := `UPDATE notes SET `
	args := []interface{}{}
	argPosition := 1

	if note.Title != "" {
		query += fmt.Sprintf("title = $%d, ", argPosition)
		args = append(args, note.Title)
		argPosition++
	}

	if note.Description != nil {
		query += fmt.Sprintf("description = $%d, ", argPosition)
		args = append(args, note.Description)
		argPosition++
	}

	if note.FilePath != "" {
		query += fmt.Sprintf("file_path = $%d, ", argPosition)
		args = append(args, note.FilePath)
		argPosition++
	}

	if note.UserID != "" {
		query += fmt.Sprintf("user_id = $%d, ", argPosition)
		args = append(args, note.UserID)
		argPosition++
	}

	if note.UniversityID != "" {
		query += fmt.Sprintf("university_id = $%d, ", argPosition)
		args = append(args, note.UniversityID)
		argPosition++
	}

	if note.CourseCode != nil { // CourseCode bir pointer olduğu için nil kontrolü yapılmalı
		query += fmt.Sprintf("course_code = $%d, ", argPosition)
		args = append(args, note.CourseCode)
		argPosition++
	}

	if note.Status != "" {
		query += fmt.Sprintf("status = $%d, ", argPosition)
		args = append(args, note.Status)
		argPosition++
	}
	// Son virgülü kaldır
	query = strings.TrimSuffix(query, ", ")

	query += fmt.Sprintf(" WHERE id = $%d RETURNING id", argPosition)
	args = append(args, note.ID)

	var uuid string
	err := r.db.QueryRow(ctx, query, args...).Scan(&uuid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", ErrNoteNotFound
		}
		return "", fmt.Errorf("not güncellenirken hata oluştu: %w", err)
	}
	return uuid, nil
}

func (r *postgresNoteRepository) GetById(ctx context.Context, id string) (*domain.Note, error) {
	var note domain.Note
	// Soft-delete edilen notların gelmemesi için `deleted_at IS NULL` kontrolü eklendi.
	row := r.db.QueryRow(ctx, `SELECT id, title, description, file_path, user_id, university_id, course_code, status, download_count, average_rating, created_at, updated_at, deleted_at FROM notes WHERE id = $1 AND deleted_at IS NULL`, id)
	err := row.Scan(&note.ID, &note.Title, &note.Description, &note.FilePath, &note.UserID, &note.UniversityID, &note.CourseCode, &note.Status, &note.DownloadCount, &note.AverageRating, &note.CreatedAt, &note.UpdatedAt, &note.DeletedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoteNotFound
		}
		return nil, fmt.Errorf("not getirilirken hata oluştu: %w", err)
	}
	return &note, nil
}

func (r *postgresNoteRepository) Delete(ctx context.Context, id string) error {
	commandTag, err := r.db.Exec(ctx,
		`UPDATE notes SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL`,
		id,
	)
	if err != nil {
		return fmt.Errorf("not silinirken hata oluştu: %w", err)
	}
	// Exec komutunun kaç satırı etkilediğini kontrol ediyoruz.
	// Eğer 0 satır etkilendiyse, o ID'ye sahip bir not bulunamadı veya zaten silinmişti.
	if commandTag.RowsAffected() == 0 {
		return ErrNoteNotFound
	}
	return nil
}
