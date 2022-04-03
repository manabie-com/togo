package model

import "gorm.io/gorm"

// Task represents the task model
type Task struct {
	Base
	Content string         `json:"content"`
	Deleted gorm.DeletedAt `json:"deleted_at"`
	UserID  int            `json:"user_id"`
}
