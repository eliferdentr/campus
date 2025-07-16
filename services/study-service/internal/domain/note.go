package domain

import (
	"time"
)

// Note, sistemdeki tek bir ders notunu temsil eden ana varlıktır.
type Note struct {
	ID            string     `json:"id" db:"id"`
	Title         string     `json:"title" db:"title"`
	Description   *string    `json:"description,omitempty" db:"description"` // Boş olabilir (nullable)
	FilePath      string     `json:"filePath" db:"file_path"` // S3 veya MinIO'daki dosyanın yolu/anahtarı
	UserID        string     `json:"userId" db:"user_id"`
	UniversityID  string     `json:"universityId" db:"university_id"`
	CourseCode    *string    `json:"courseCode,omitempty" db:"course_code"` // Boş olabilir (nullable)
	Status        Status     `json:"status" db:"status"`
	DownloadCount int64      `json:"downloadCount" db:"download_count"`
	AverageRating float64    `json:"averageRating" db:"average_rating"`
	CreatedAt     time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt     time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt     *time.Time `json:"deletedAt,omitempty" db:"deleted_at"` // Soft delete için pointer olmalı
}