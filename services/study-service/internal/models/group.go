package models

import (
	"time"

	"github.com/google/uuid"
)

type Group struct {
	ID uuid.UUID  `gorm:"primaryKey"`
	Name string `gorm:"column:group_name"`
	Description string `gorm:"group_description"`
	CourseCode string `gorm:"column:course_code"`
	OwnerId uuid.UUID `gorm:"column:owner_id"`
	CreatedAat time.Time `gorm:"column:created_at"`
}