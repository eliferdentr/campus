package repository_test

import (
	"context"
	"regexp"
	"testing"
	"time"

	"study-service/internal/domain"
	r "study-service/internal/repository"

	"github.com/jackc/pgx/v5"
	pgxmock "github.com/pashagolub/pgxmock/v3"
	assert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// func StartTests(t *testing.T) {
// 	repo := r.NewPostgresNoteRepository()

// 	uuid :=testPostgresCreate(&note, t, repo)
// 	note.Title = "Test2 Note"
// 	uuid = testPostgresUpdate(&note, t, repo)
// 	updatedNote := testPostgresGetById(uuid, t, repo)
// 	assert.Equal(t, "Test2 Note", updatedNote.Title)
// 	testPostgresDelete(uuid, t, repo)
// }

func TestPostgresNoteRepository_Create(t *testing.T) {
	note := &domain.Note{
		Title:        "Test Note",
		Description:  nil,
		FilePath:     "/path/to/file",
		UserID:       "user-123",
		UniversityID: "uni-456",
		CourseCode:   nil,
		Status:       domain.StatusPending,
	}

	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()
	repo := r.NewPostgresNoteRepository(mock)
	expectedUUID := "test-uuid-123"

	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO notes (title, description, file_path, user_id, university_id, course_code) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`)).
		WithArgs(note.Title, note.Description, note.FilePath, note.UserID, note.UniversityID, note.CourseCode).
		WillReturnRows(pgxmock.NewRows([]string{"id"}).AddRow(expectedUUID))

	uuid, err := repo.Create(context.Background(), note)

	assert.NoError(t, err)
	assert.Equal(t, expectedUUID, uuid)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPostgresNoteRepository_Update_Success(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := r.NewPostgresNoteRepository(mock)

	noteToUpdate := &domain.Note{
		ID:    "test-uuid-123",
		Title: "Yeni Başlık",
	}

	expectedSQL := regexp.QuoteMeta(`UPDATE notes SET title = $1 WHERE id = $2 RETURNING id`)

	// Beklenen argümanlar
	expectedArgs := []interface{}{noteToUpdate.Title, noteToUpdate.ID}

	mock.ExpectQuery(expectedSQL).
		WithArgs(expectedArgs...). // WithArgs, birebir eşleşme bekler.
		WillReturnRows(pgxmock.NewRows([]string{"id"}).AddRow(noteToUpdate.ID))

	// Metodu çağır
	updatedID, err := repo.Update(context.Background(), noteToUpdate)

	assert.NoError(t, err)
	assert.Equal(t, noteToUpdate.ID, updatedID)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPostgresNoteRepository_Update_NotFound(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := r.NewPostgresNoteRepository(mock)

	noteToUpdate := &domain.Note{
		ID:    "yok-boyle-bir-uuid",
		Title: "Yeni Başlık",
	}

	// Beklenen sorgu ve argümanlar aynı
	expectedSQL := regexp.QuoteMeta(`UPDATE notes SET title = $1 WHERE id = $2 RETURNING id`)
	expectedArgs := []interface{}{noteToUpdate.Title, noteToUpdate.ID}

	// ANAHTAR NOKTA: Bu sefer bir satır döndürmek yerine hata döndürmesini söylüyoruz.
	mock.ExpectQuery(expectedSQL).
		WithArgs(expectedArgs...).
		WillReturnError(pgx.ErrNoRows)

	// Metodu çağır
	_, err = repo.Update(context.Background(), noteToUpdate)

	// Sonuçları kontrol et
	// Bu sefer bir hata bekliyoruz.
	assert.Error(t, err)

	// Ve dönen hatanın bizim özel 'NotFound' hatamız olduğunu doğruluyoruz.
	// Bu, repository'nin hata yönetiminin doğru çalıştığını test eder.
	assert.ErrorIs(t, err, r.ErrNoteNotFound)

	// Mock beklentilerinin karşılandığından emin ol
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPostgresNoteRepository_GetById_Success(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := r.NewPostgresNoteRepository(mock)

	expectedNote := &domain.Note{
		ID:           "test-uuid-123",
		Title:        "Test Note",
		Description:  nil,
		FilePath:     "/path/to/file",
		UserID:       "user-123",
		UniversityID: "uni-456",
		CourseCode:   nil,
		Status:       domain.StatusPending,
	}
	rows := pgxmock.NewRows([]string{
		"id", "title", "description", "file_path", "user_id", "university_id", "course_code",
		"status", "download_count", "average_rating", "created_at", "updated_at", "deleted_at",
	}).AddRow(
		expectedNote.ID, expectedNote.Title, expectedNote.Description, expectedNote.FilePath, expectedNote.UserID, expectedNote.UniversityID,
		expectedNote.CourseCode, expectedNote.Status, expectedNote.DownloadCount, expectedNote.AverageRating, time.Now(), time.Now(), time.Now(),
	)

    mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, title, description, file_path, user_id, university_id, course_code, status, download_count, average_rating, created_at, updated_at, deleted_at FROM notes WHERE id = $1 AND deleted_at IS NULL`)).
        WithArgs(expectedNote.ID).
        WillReturnRows(rows)
		
    foundNote, err := repo.GetById(context.Background(), expectedNote.ID)


    assert.NoError(t, err)
    require.NotNil(t, foundNote)
    assert.Equal(t, expectedNote.ID, foundNote.ID)
    assert.Equal(t, expectedNote.Title, foundNote.Title)

    assert.NoError(t, mock.ExpectationsWereMet())
} 

func TestPostgresNoteRepository_GetById_NotFound(t *testing.T) {
mock, err := pgxmock.NewPool()
    require.NoError(t, err)
    defer mock.Close()

    repo := r.NewPostgresNoteRepository(mock)

    nonExistentID := "yok-boyle-bir-uuid"

	expectedSQL := `SELECT id, title, description, file_path, user_id, university_id, course_code, status, download_count, average_rating, created_at, updated_at, deleted_at FROM notes WHERE id = $1 AND deleted_at IS NULL`
    // Mock'a "satır bulunamadı" hatası döndürmesini söylüyoruz.
    mock.ExpectQuery(regexp.QuoteMeta(expectedSQL)).
        WithArgs(nonExistentID).
        WillReturnError(pgx.ErrNoRows)

    // Metodu çağırıyoruz
    foundNote, err := repo.GetById(context.Background(), nonExistentID)

    // Sonuçları kontrol ediyoruz
    assert.Error(t, err) // Hata bekliyoruz
    assert.Nil(t, foundNote) // Not'un nil olmasını bekliyoruz
    
    // Dönen hatanın bizim özel ErrNoteNotFound hatamız olduğunu doğruluyoruz.
    assert.ErrorIs(t, err, r.ErrNoteNotFound)

    assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPostgresNoteRepository_Delete_NotFound(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := r.NewPostgresNoteRepository(mock)
	idToDelete := "yok-boyle-bir-uuid"

	// Metod Exec kullandığı için mock'a ExpectExec diyoruz.
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE notes SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL`)).
		WithArgs(idToDelete).
		// ANAHTAR NOKTA: Bu sefer 0 satırın etkilendiğini söylüyoruz.
		WillReturnResult(pgxmock.NewResult("UPDATE", 0))

	// Metodu çağırıyoruz
	err = repo.Delete(context.Background(), idToDelete)

	// Sonuçları kontrol ediyoruz
	assert.Error(t, err) // Hata bekliyoruz.
	assert.ErrorIs(t, err, r.ErrNoteNotFound) // Özel 'NotFound' hatası olmalı.

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPostgresNoteRepository_Delete_Success(t *testing.T) {
mock, err := pgxmock.NewPool()
    require.NoError(t, err)
    defer mock.Close()

    repo := r.NewPostgresNoteRepository(mock)
    idToDelete := "test-uuid-123"

    // Metod Exec kullandığı için mock'a ExpectExec diyoruz.
    mock.ExpectExec(regexp.QuoteMeta(`UPDATE notes SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL`)).
        WithArgs(idToDelete).
        // Exec için WillReturnRows yerine WillReturnResult kullanırız.
        // "UPDATE" komutunun 1 satırı etkilediğini simüle ediyoruz.
        WillReturnResult(pgxmock.NewResult("UPDATE", 1))

    // Metodu çağırıyoruz
    err = repo.Delete(context.Background(), idToDelete)

    // Sonuçları kontrol ediyoruz
    assert.NoError(t, err) // Hata olmamasını bekliyoruz.

    assert.NoError(t, mock.ExpectationsWereMet())
}
