package task

import (
	"gorm.io/gorm"
)

// Task represents task service
type Task struct {
	db *gorm.DB
}

// New creates new task service
func New(db *gorm.DB) *Task {
	return &Task{
		db: db,
	}
}
