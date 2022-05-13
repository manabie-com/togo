package service

import (
	"context"
	"sync"

	domain "github.com/nvhai245/togo/internal/domain"
)

type taskService struct {
	taskRepository domain.TaskRepository
}

var (
	instance *taskService
	once     sync.Once
)

// NewTaskService injects the task repository
func NewTaskService(r domain.TaskRepository) domain.TaskService {
	once.Do(func() {
		instance = &taskService{taskRepository: r}
	})
	return instance
}

func (s *taskService) CreateTask(ctx context.Context, userID int64, content string) (*domain.Task, error) {
	err := s.taskRepository.CheckTaskByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return s.taskRepository.CreateTask(ctx, userID, content)
}

func (s *taskService) GetAllTaskByUserID(ctx context.Context, userID int64) ([]*domain.Task, error) {
	return s.taskRepository.GetAllTaskByUserID(ctx, userID)
}
