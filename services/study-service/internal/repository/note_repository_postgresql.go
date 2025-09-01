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

func (r *postgresNoteRepository) Create(ctx context.Context, note *domain.Note) (*domain.Note, error) {
    var createdNote domain.Note 
    query := `
        INSERT INTO notes (title, description, file_path, user_id, university_id, course_code, status) 
        VALUES ($1, $2, $3, $4, $5, $6, $7) 
        RETURNING *`

    row := r.db.QueryRow(ctx, query,
        note.Title, note.Description, note.FilePath, note.UserID, note.UniversityID, note.CourseCode, note.Status)


    err := row.Scan(
        &createdNote.ID,
        &createdNote.Title,
        &createdNote.Description,
        &createdNote.FilePath,
        &createdNote.UserID,
        &createdNote.UniversityID,
        &createdNote.CourseCode,
        &createdNote.Status,
        &createdNote.DownloadCount,
        &createdNote.AverageRating,
        &createdNote.CreatedAt,
        &createdNote.UpdatedAt,
        &createdNote.DeletedAt,
    )

    if err != nil {
        return nil, fmt.Errorf("not oluşturulurken veritabanı hatası oluştu: %w", err)
    }
    return &createdNote, nil
}


func (r *postgresNoteRepository) Update(ctx context.Context, note *domain.Note) (*domain.Note, error) {
    clauses := []string{}
    args := []interface{}{}
    argPosition := 1

    if note.Title != "" {
        clauses = append(clauses, fmt.Sprintf("title = $%d", argPosition))
        args = append(args, note.Title)
        argPosition++
    }
    if note.Description != nil {
        clauses = append(clauses, fmt.Sprintf("description = $%d", argPosition))
        args = append(args, note.Description)
        argPosition++
    }

	if note.FilePath != "" {
		clauses = append(clauses, fmt.Sprintf("file_path = $%d", argPosition))
		args = append(args, note.FilePath)
		argPosition++
	}

	if note.UserID != "" {
		clauses = append(clauses, fmt.Sprintf("user_id = $%d", argPosition))
		args = append(args, note.UserID)
		argPosition++
	}

	if note.UniversityID != "" {
		clauses = append(clauses, fmt.Sprintf("university_id = $%d", argPosition))
		args = append(args, note.UniversityID)
		argPosition++
	}

	if note.CourseCode != nil { // CourseCode bir pointer olduğu için nil kontrolü yapılmalı
		clauses = append(clauses, fmt.Sprintf("course_code = $%d", argPosition))
		args = append(args, note.CourseCode)
		argPosition++
	}

	if note.Status != "" {
		clauses = append(clauses, fmt.Sprintf("status = $%d", argPosition))
		args = append(args, note.Status)
		argPosition++
	}

    if len(clauses) == 0 {
        // Güncellenecek bir şey yoksa, hata döndürmek yerine mevcut notu döndürebiliriz.
        // Ama önce onu veritabanından çekmemiz gerekir. Şimdilik hata döndürmek daha basit.
        return nil, errors.New("güncellenecek alan bulunamadı")
    }

    clauses = append(clauses, "updated_at = NOW()")
    setClause := strings.Join(clauses, ", ")


    query := fmt.Sprintf("UPDATE notes SET %s WHERE id = $%d RETURNING *",
        setClause,
        argPosition,
    )
    args = append(args, note.ID)
    var updatedNote domain.Note
    err := r.db.QueryRow(ctx, query, args...).Scan(
        &updatedNote.ID,
        &updatedNote.Title,
        &updatedNote.Description,
        &updatedNote.FilePath,
        &updatedNote.UserID,
        &updatedNote.UniversityID,
        &updatedNote.CourseCode,
        &updatedNote.Status,
        &updatedNote.DownloadCount,
        &updatedNote.AverageRating,
        &updatedNote.CreatedAt,
        &updatedNote.UpdatedAt,
        &updatedNote.DeletedAt, 
    )

    if err != nil {
        if errors.Is(err, pgx.ErrNoRows) {
            return nil, ErrNoteNotFound
        }
        return nil, fmt.Errorf("not güncellenirken veritabanı hatası oluştu: %w", err)
    }
    return &updatedNote, nil
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
