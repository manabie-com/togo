package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Base model for common columns
type Base struct {
	UserID      uuid.UUID  `json:"userId" gorm:"column:user_id;"`
	IsActive    bool       `json:"isActive" gorm:"column:is_active;"`
	CreatedBy   uuid.UUID  `json:"createdBy" gorm:"column:created_by;"`
	UpdatedBy   *uuid.UUID `json:"updatedBy" gorm:"column:updated_by;"`
	CreatedWhen time.Time  `json:"createdWhen" gorm:"column:created_when;"`
	UpdatedWhen *time.Time `json:"updatedWhen" gorm:"column:updated_when;"`
}
