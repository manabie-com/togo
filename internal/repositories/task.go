package repositories

import (
	"gorm.io/gorm"
)

type TaskRepository interface {
}

func newTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

type taskRepository struct{ db *gorm.DB }
