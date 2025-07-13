package service

import (
	"context"
	"time"

	"campus-project.com/study-service/internal/domain"
	"campus-project.com/study-service/internal/repository"
	"github.com/google/uuid"
)

// NoteService, notlarla ilgili iş mantığı için arayüzü tanımlar.
type NoteService interface {
	Create(ctx context.Context, userID, title, description, courseCode string) (*domain.Note, error)
}

type noteService struct {
	repo repository.NoteRepository
}

// NewNoteService, noteService için yeni bir örnek oluşturur.
func NewNoteService(repo repository.NoteRepository) NoteService {
	return &noteService{repo: repo}
}

func (s *noteService) Create(ctx context.Context, userID, title, description, courseCode string) (*domain.Note, error) {
	note := &domain.Note{
		ID:          uuid.NewString(),
		Title:       title,
		Description: description,
		UserID:      userID,
		CourseCode:  courseCode,
		CreatedAt:   time.Now(),
	}

	if err := s.repo.Save(ctx, note); err != nil {
		return nil, err
	}

	return note, nil
}
