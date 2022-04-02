package model

import (
	"time"
)

// Base contains common fields for all models
// Do not use gorm.Model because of uint ID
type Base struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
