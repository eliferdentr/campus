package repository

import (
	"context"
	"study-service/internal/domain" // Projenizin go.mod dosyasındaki module adının 'study-service' olduğu varsayılmıştır.
)

// NoteRepository, not verileri üzerindeki operasyonlar için arayüzü tanımlar.
type NoteRepository interface {
	// Save, verilen notu veritabanına kaydeder.
	Create(ctx context.Context, note *domain.Note) (*domain.Note, error)
	GetById(ctx context.Context, id string) (*domain.Note, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, note *domain.Note) (*domain.Note, error)
	
}
