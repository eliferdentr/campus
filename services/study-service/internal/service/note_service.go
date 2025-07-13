package service

import (
	"context"
	"study-service/internal/domain"
	"study-service/internal/repository"
	"time"

	"github.com/google/uuid"
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
	note := &domain.Note{
		ID:          uuid.NewString(),
		Title:       title,
		Description: description,
		UserID:      userID,
		CourseCode:  courseCode,
		CreatedAt:   time.Now().UTC(),
		// FilePath ve UniversityID gibi alanlar daha sonraki adımlarda atanabilir.
	}

	if err := s.noteRepo.Save(ctx, note); err != nil {
		return nil, err
	}

	return note, nil
}
