package models

import (
	uuid "github.com/satori/go.uuid"
)

// Task model for `tasks` table
type Task struct {
	Base
	TaskID      uuid.UUID `json:"taskId" gorm:"column:task_id;"`
	Title       string    `json:"title" gorm:"column:title;"`
	Description *string   `json:"description" gorm:"column:description;"`
}
