package domain

import "time"

// Note represents a user-uploaded note.
type Note struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	FilePath     string    `json:"filePath"`
	UserID       string    `json:"userId"`
	CourseCode   string    `json:"courseCode"`
	UniversityID string    `json:"universityId"`
	CreatedAt    time.Time `json:"createdAt"`
}
