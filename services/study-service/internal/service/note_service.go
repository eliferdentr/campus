package service

import (
	"context"
	"errors"
	"study-service/internal/domain"
)

var (
	ErrNoteNotFound = errors.New("not bulunamadı")
)

type NoteService interface {
	Create(ctx context.Context, note domain.Note) (*domain.Note, error)
	Update(ctx context.Context, note domain.Note) (*domain.Note, error)
	GetById(ctx context.Context, uuid string) (*domain.Note, error)
	Delete(ctx context.Context, uuid string) error
}