package request

import (
	"time"
	"github.com/google/uuid"
)

type CreateLiveSessionRequest struct {
	GroupID     uuid.UUID `json:"group_id" binding:"required"`
	Platform    string    `json:"platform" binding:"required"`     // "zoom" veya "meet"
	SessionLink string    `json:"session_link" binding:"required"` // Zoom/Meet URL
	StartTime   time.Time `json:"start_time" binding:"required"`
	DurationMin int       `json:"duration_min,omitempty"` 
}