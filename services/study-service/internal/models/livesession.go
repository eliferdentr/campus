package models

import (
	"time"
	"github.com/google/uuid"
)

type LiveSession struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	GroupID     uuid.UUID `gorm:"column:group_id"`
	Platform    string    `gorm:"column:platform_id"`
	SessionLink string    `gorm:"column:session_link"`
	StartTime   time.Time `gorm:"column:start_time"`
	DurationMin int       `gorm:"column:duration_min"`
	CreatedAt   time.Time `gorm:"column:created_at"`
}