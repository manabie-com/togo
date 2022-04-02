package model

import (
	"time"
)

// Base contains common fields for all models
// Do not use gorm.Model because of uint ID
type Base struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
}
