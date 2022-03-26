package storage

import (
	"context"

	"github.com/luongdn/togo/models"
	"gorm.io/gorm"
)

func NewTaskStore(db *gorm.DB) *taskStore {
	return &taskStore{
		sqlDB: db,
	}
}

type taskStore struct {
	sqlDB *gorm.DB
}

func (s *taskStore) ListTasks(ctx context.Context, user_id string) ([]models.Task, error) {
	tasks := []models.Task{}
	result := s.sqlDB.Where("user_id = ?", user_id).Find(&tasks)

	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (s *taskStore) CreateTask(ctx context.Context, user_id string, task *models.Task) error {
	task.UserID = user_id
	result := s.sqlDB.Create(task)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
