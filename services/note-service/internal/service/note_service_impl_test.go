package service_test

import (
	"context"
	"errors"
	"study-service/internal/domain"
	"study-service/internal/repository"
	"study-service/internal/service"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockNoteRepository, NoteRepository arayüzü için bir mock'tur.
type MockNoteRepository struct {
	mock.Mock
}

func (m *MockNoteRepository) Create(ctx context.Context, note *domain.Note) (*domain.Note, error) {
	args := m.Called(ctx, note)
	if ret := args.Get(0); ret != nil {
		return ret.(*domain.Note), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockNoteRepository) Update(ctx context.Context, note *domain.Note) (*domain.Note, error) {
	args := m.Called(ctx, note)
	if ret := args.Get(0); ret != nil {
		return ret.(*domain.Note), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockNoteRepository) GetById(ctx context.Context, id string) (*domain.Note, error) {
	args := m.Called(ctx, id)
	if ret := args.Get(0); ret != nil {
		return ret.(*domain.Note), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockNoteRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestNoteService_Create(t *testing.T) {
	tests := []struct {
		name          string
		noteToCreate  domain.Note
		setupMock     func(repo *MockNoteRepository)
		checkResult   func(t *testing.T, got *domain.Note, want *domain.Note)
		wantErr       bool
		expectedError string
	}{
		{
			name: "Başarılı oluşturma",
			noteToCreate: domain.Note{
				Title:  "Geçerli Başlık",
				UserID: "user-123",
			},
			setupMock: func(repo *MockNoteRepository) {
				expectedNoteArg := &domain.Note{
					Title:         "Geçerli Başlık",
					UserID:        "user-123",
					Status:        domain.StatusPending,
					DownloadCount: 0,
					AverageRating: 0.0,
				}
				returnedNote := &domain.Note{
					ID:            "new-uuid",
					Title:         "Geçerli Başlık",
					UserID:        "user-123",
					Status:        domain.StatusPending,
					CreatedAt:     time.Now(),
					UpdatedAt:     time.Now(),
				}
				repo.On("Create", mock.Anything, expectedNoteArg).Return(returnedNote, nil).Once()
			},
			checkResult: func(t *testing.T, got *domain.Note, want *domain.Note) {
				assert.Equal(t, want.ID, got.ID)
				assert.Equal(t, want.Title, got.Title)
				assert.Equal(t, want.UserID, got.UserID)
				assert.Equal(t, want.Status, got.Status)
			},
			wantErr: false,
		},
		{
			name: "Doğrulama Hatası - Boş Başlık",
			noteToCreate: domain.Note{
				UserID: "user-123",
			},
			setupMock:     func(repo *MockNoteRepository) {},
			wantErr:       true,
			expectedError: "başlık boş olamaz",
		},
		{
			name: "Doğrulama Hatası - Boş UserID",
			noteToCreate: domain.Note{
				Title: "Geçerli Başlık",
			},
			setupMock:     func(repo *MockNoteRepository) {},
			wantErr:       true,
			expectedError: "kullanıcı boş olamaz",
		},
		{
			name: "Repository Hatası",
			noteToCreate: domain.Note{
				Title:  "Geçerli Başlık",
				UserID: "user-123",
			},
			setupMock: func(repo *MockNoteRepository) {
				expectedNoteArg := &domain.Note{
					Title:         "Geçerli Başlık",
					UserID:        "user-123",
					Status:        domain.StatusPending,
					DownloadCount: 0,
					AverageRating: 0.0,
				}
				repo.On("Create", mock.Anything, expectedNoteArg).Return(nil, errors.New("database error")).Once()
			},
			wantErr:       true,
			expectedError: "database error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockNoteRepository)
			noteSvc := service.NewNoteService(mockRepo)
			tt.setupMock(mockRepo)

			got, err := noteSvc.Create(context.Background(), tt.noteToCreate)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.expectedError != "" {
					assert.EqualError(t, err, tt.expectedError)
				}
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
				if tt.checkResult != nil {
					// Sadece ID'yi ve temel alanları kontrol etmek için 'want' nesnesini hazırlıyoruz
					want := &domain.Note{ID: got.ID, Title: tt.noteToCreate.Title, UserID: tt.noteToCreate.UserID, Status: domain.StatusPending}
					tt.checkResult(t, got, want)
				}
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestNoteService_Update(t *testing.T) {
	existingDesc := "Orijinal Açıklama"
	existingNote := &domain.Note{
		ID:            "note-123",
		Title:         "Orijinal Başlık",
		Description:   &existingDesc,
		UserID:        "user-abc",
		Status:        domain.StatusPending,
		DownloadCount: 50,
	}

	tests := []struct {
		name          string
		partialUpdate domain.Note
		setupMock     func(repo *MockNoteRepository)
		checkResult   func(t *testing.T, got *domain.Note)
		wantErr       bool
		expectedError string
	}{
		{
			name: "Başarılı - Başlık güncelleme",
			partialUpdate: domain.Note{
				ID:    "note-123",
				Title: "Güncellenmiş Başlık",
			},
			setupMock: func(repo *MockNoteRepository) {
				repo.On("GetById", mock.Anything, "note-123").Return(existingNote, nil).Once()
				repo.On("Update", mock.Anything, mock.MatchedBy(func(n *domain.Note) bool {
					return n.ID == "note-123" && n.Title == "Güncellenmiş Başlık" && !n.UpdatedAt.IsZero()
				})).Return(&domain.Note{ID: "note-123", Title: "Güncellenmiş Başlık", Description: &existingDesc}, nil).Once()
			},
			checkResult: func(t *testing.T, got *domain.Note) {
				assert.Equal(t, "Güncellenmiş Başlık", got.Title)
				assert.Equal(t, existingDesc, *got.Description)
			},
			wantErr: false,
		},
		{
			name: "Başarılı - Sıfır değerli güncellemeleri yoksayar",
			partialUpdate: domain.Note{
				ID:            "note-123",
				Title:         "Yeni Başlık",
				DownloadCount: 0, // Bu yoksayılmalı
			},
			setupMock: func(repo *MockNoteRepository) {
				repo.On("GetById", mock.Anything, "note-123").Return(existingNote, nil).Once()
				updatedNote := *existingNote
				updatedNote.Title = "Yeni Başlık"
				repo.On("Update", mock.Anything, mock.MatchedBy(func(n *domain.Note) bool {
					return n.ID == "note-123" && n.Title == "Yeni Başlık" && n.DownloadCount == 50
				})).Return(&updatedNote, nil).Once()
			},
			checkResult: func(t *testing.T, got *domain.Note) {
				assert.Equal(t, "Yeni Başlık", got.Title)
				assert.Equal(t, int64(50), got.DownloadCount)
			},
			wantErr: false,
		},
		{
			name:          "Doğrulama Hatası - ID yok",
			partialUpdate: domain.Note{Title: "Bir Başlık"},
			setupMock:     func(repo *MockNoteRepository) {},
			wantErr:       true,
			expectedError: "güncelleme için not ID'si boş olamaz",
		},
		{
			name:          "GetById'den Bulunamadı Hatası",
			partialUpdate: domain.Note{ID: "not-found-id"},
			setupMock: func(repo *MockNoteRepository) {
				repo.On("GetById", mock.Anything, "not-found-id").Return(nil, repository.ErrNoteNotFound).Once()
			},
			wantErr:       true,
			expectedError: service.ErrNoteNotFound.Error(),
		},
		{
			name: "Update üzerinde Repository Hatası",
			partialUpdate: domain.Note{
				ID:    "note-123",
				Title: "Güncellenmiş Başlık",
			},
			setupMock: func(repo *MockNoteRepository) {
				repo.On("GetById", mock.Anything, "note-123").Return(existingNote, nil).Once()
				repo.On("Update", mock.Anything, mock.Anything).Return(nil, errors.New("db update failed")).Once()
			},
			wantErr:       true,
			expectedError: "db update failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockNoteRepository)
			noteSvc := service.NewNoteService(mockRepo)
			tt.setupMock(mockRepo)

			got, err := noteSvc.Update(context.Background(), tt.partialUpdate)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.expectedError != "" {
					// Not: Bu test, Update fonksiyonunun GetById'den gelen ErrNoteNotFound hatasını
					// sarmalayıp service.ErrNoteNotFound'a dönüştürdüğünü varsayar.
					// Mevcut implementasyon bunu yapmıyorsa bu test başarısız olur ve bu bir iyileştirme alanıdır.
					assert.EqualError(t, err, tt.expectedError)
				}
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
				if tt.checkResult != nil {
					tt.checkResult(t, got)
				}
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestNoteService_GetById(t *testing.T) {
	expectedNote := &domain.Note{ID: "note-123", Title: "Test Notu"}

	tests := []struct {
		name          string
		idToGet       string
		setupMock     func(repo *MockNoteRepository)
		want          *domain.Note
		wantErr       bool
		expectedError error
	}{
		{
			name:    "Başarılı getirme",
			idToGet: "note-123",
			setupMock: func(repo *MockNoteRepository) {
				repo.On("GetById", mock.Anything, "note-123").Return(expectedNote, nil).Once()
			},
			want:    expectedNote,
			wantErr: false,
		},
		{
			name:          "Doğrulama Hatası - Boş ID",
			idToGet:       "",
			setupMock:     func(repo *MockNoteRepository) {},
			wantErr:       true,
			expectedError: errors.New("uuid boş olamaz"),
		},
		{
			name:    "Bulunamadı Hatası",
			idToGet: "not-found-id",
			setupMock: func(repo *MockNoteRepository) {
				repo.On("GetById", mock.Anything, "not-found-id").Return(nil, repository.ErrNoteNotFound).Once()
			},
			wantErr:       true,
			expectedError: service.ErrNoteNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockNoteRepository)
			noteSvc := service.NewNoteService(mockRepo)
			tt.setupMock(mockRepo)

			got, err := noteSvc.GetById(context.Background(), tt.idToGet)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.expectedError != nil {
					if errors.Is(err, service.ErrNoteNotFound) {
						assert.ErrorIs(t, err, service.ErrNoteNotFound)
					} else {
						assert.EqualError(t, err, tt.expectedError.Error())
					}
				}
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestNoteService_Delete(t *testing.T) {
	tests := []struct {
		name          string
		idToDelete    string
		setupMock     func(repo *MockNoteRepository)
		wantErr       bool
		expectedError error
	}{
		{
			name:       "Başarılı silme",
			idToDelete: "note-123",
			setupMock: func(repo *MockNoteRepository) {
				repo.On("Delete", mock.Anything, "note-123").Return(nil).Once()
			},
			wantErr: false,
		},
		{
			name:          "Doğrulama Hatası - Boş ID",
			idToDelete:    "",
			setupMock:     func(repo *MockNoteRepository) {},
			wantErr:       true,
			expectedError: errors.New("uuid boş olamaz"),
		},
		{
			name:       "Bulunamadı Hatası",
			idToDelete: "not-found-id",
			setupMock: func(repo *MockNoteRepository) {
				repo.On("Delete", mock.Anything, "not-found-id").Return(repository.ErrNoteNotFound).Once()
			},
			wantErr:       true,
			expectedError: service.ErrNoteNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockNoteRepository)
			noteSvc := service.NewNoteService(mockRepo)
			tt.setupMock(mockRepo)

			err := noteSvc.Delete(context.Background(), tt.idToDelete)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.expectedError != nil {
					if errors.Is(err, service.ErrNoteNotFound) {
						assert.ErrorIs(t, err, service.ErrNoteNotFound)
					} else {
						assert.EqualError(t, err, tt.expectedError.Error())
					}
				}
			} else {
				assert.NoError(t, err)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

