package models

import (
	"time"
	"github.com/google/uuid"
)

type Platform struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	GroupID     uuid.UUID `gorm:"column:group_id"`
	Name    string    `gorm:"column:platform_name"`
	SessionLink string    `gorm:"column:session_link"`
	StartTime   time.Time `gorm:"column:start_time"`
	DurationMin int       `gorm:"column:duration_min"`
	CreatedAt   time.Time `gorm:"column:created_at"`
}