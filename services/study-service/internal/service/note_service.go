package service

import (
	"context"
	"study-service/internal/domain"
	"study-service/internal/repository"
	"time"
)

// NoteService, notlarla ilgili iş mantığı operasyonlarını tanımlar.
type NoteService interface {
	Create(ctx context.Context, userID, title, description, courseCode string) (*domain.Note, error)
}

// noteService, NoteService arayüzünün implementasyonudur.
type noteService struct {
	noteRepo repository.NoteRepository
}

// NewNoteService, yeni bir NoteService örneği oluşturur.
func NewNoteService(repo repository.NoteRepository) NoteService {
	return &noteService{noteRepo: repo}
}

// Create, yeni bir not nesnesi oluşturur ve repository aracılığıyla kaydeder.
func (s *noteService) Create(ctx context.Context, userID, title, description, courseCode string) (*domain.Note, error) {
	// Description ve CourseCode alanları veritabanında NULL olabildiği için
	// domain.Note struct'ında pointer (*string) olarak tanımlanmışlar.
	// Gelen string boş ise nil, değilse string'in adresini atıyoruz.
	var descPtr *string
	if description != "" {
		descPtr = &description
	}

	var courseCodePtr *string
	if courseCode != "" {
		courseCodePtr = &courseCode
	}

	note := &domain.Note{
		// ID, repository katmanında veritabanı tarafından oluşturulur (RETURNING id).
		// Bu yüzden service katmanında ID ataması yapmıyoruz.
		Title:        title,
		Description:  descPtr,
		UserID:       userID,
		CourseCode:   courseCodePtr,
		CreatedAt:    time.Now().UTC(),
		// FilePath ve UniversityID gibi alanlar daha sonraki adımlarda atanabilir.
	}

	createdID, err := s.noteRepo.Create(ctx, note)
	if err != nil {
		return nil, err
	}
	note.ID = createdID
	return note, nil
}
