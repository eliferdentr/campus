package service

import (
	"context"
	"errors"
	"study-service/internal/domain"
	"study-service/internal/repository"
	"time"
)

type noteService struct {
	noteRepo repository.NoteRepository
}

func NewNoteService(repo repository.NoteRepository) NoteService {
	return &noteService{noteRepo: repo}
}

func (s *noteService) Create(ctx context.Context, note domain.Note) (*domain.Note, error) {
	if note.Title == "" {
		return nil, errors.New("başlık boş olamaz")
	}
	if note.UserID == "" {
		return nil, errors.New("kullanıcı boş olamaz")
	}
	note.Status = domain.StatusPending // Varsayılan durumu ata
	note.DownloadCount = 0
	note.AverageRating = 0.0
	createdNote, err := s.noteRepo.Create(ctx, &note)
	if err != nil {
		return nil, err
	}
	return createdNote, nil

}
func (s *noteService) Update(ctx context.Context, partialUpdate domain.Note) (*domain.Note, error) {
	// 1. Validasyon: Güncelleme için not ID'si zorunludur.
	if partialUpdate.ID == "" {
		return nil, errors.New("güncelleme için not ID'si boş olamaz")
	}

	// 2. GETİR: Güncellenecek kaydın tam ve mevcut halini veritabanından al.
	existingNote, err := s.noteRepo.GetById(ctx, partialUpdate.ID)
	if err != nil {
		// Not bulunamadıysa veya başka bir veritabanı hatası varsa işlemi burada sonlandır.
		return nil, err
	}
	//Yetki Kontrolü (Authorization)
	// requestingUserID := getUserFromContext(ctx) // Gerçek uygulamada JWT'den vs. gelir.
	// if existingNote.UserID != requestingUserID {
	//     return nil, errors.New("bu notu güncelleme yetkiniz yok")
	// }

	// if existingNote.Status == domain.StatusApproved {
	//     return nil, errors.New("onaylanmış notlar güncellenemez")
	// }
	// ---

	if partialUpdate.Title != "" {
		existingNote.Title = partialUpdate.Title
	}
	if partialUpdate.Description != nil {
		existingNote.Description = partialUpdate.Description
	}
	if partialUpdate.FilePath != "" {
		existingNote.FilePath = partialUpdate.FilePath
	}
	if partialUpdate.Status != "" {
		existingNote.Status = partialUpdate.Status
	}
	if partialUpdate.CourseCode != nil {
		existingNote.CourseCode = partialUpdate.CourseCode
	}

	existingNote.UpdatedAt = time.Now()
	return s.noteRepo.Update(ctx, existingNote)
}
func (s *noteService) GetById(ctx context.Context, uuid string) (*domain.Note, error) {
	if uuid == "" {
		return nil, errors.New("uuid boş olamaz")
	}
	note, err := s.noteRepo.GetById(ctx, uuid)
	if err != nil {
		if errors.Is(err, repository.ErrNoteNotFound) {
			return nil, ErrNoteNotFound
		}
		return nil, err
	}
	return note, nil

}
func (s *noteService) Delete(ctx context.Context, uuid string) error {
	if uuid == "" {
		return errors.New("uuid boş olamaz")
	}
	err := s.noteRepo.Delete(ctx, uuid)
	if err != nil {
		if errors.Is(err, repository.ErrNoteNotFound) {
			return ErrNoteNotFound
		}
		return err
	}
	return nil
}
