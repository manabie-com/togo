package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// User model for users table
type User struct {
	UserID         uuid.UUID  `json:"id" gorm:"column:user_id;primary_key;"`
	Username       string     `json:"username" gorm:"column:username;"`
	TaskDailyLimit int32      `json:"taskDailyLimit" gorm:"column:task_daily_limit;"`
	IsActive       bool       `json:"isActive" gorm:"column:is_active;"`
	CreatedWhen    time.Time  `json:"createdWhen" gorm:"column:created_when;"`
	UpdatedWhen    *time.Time `json:"updatedWhen" gorm:"column:updated_when;"`
	CreatedBy      uuid.UUID  `json:"createdBy" gorm:"column:created_by;"`
	UpdatedBy      *uuid.UUID `json:"updatedBy" gorm:"column:updated_by;"`
}
