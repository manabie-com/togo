package services

import (
	"context"
	"togo/internal/domain"
	"togo/internal/repository"
)

type taskService struct {
	userRepo      repository.UserRepository
	taskRepo      repository.TaskRepository
	taskLimitRepo repository.TaskLimitRepository
}

// NewTaskService service constructor
func NewTaskService(
	userRepo repository.UserRepository,
	taskRepo repository.TaskRepository,
	taskLimitRepo repository.TaskLimitRepository,
) domain.TaskService {
	return &taskService{
		userRepo,
		taskRepo,
		taskLimitRepo,
	}
}

func (s taskService) Create(ctx context.Context, task *domain.Task) (*domain.Task, error) {

	return nil, nil
}
func (s taskService) UpdateByID(ctx context.Context, id uint, update *domain.Task) (*domain.Task, error) {
	return nil, nil
}
func (s taskService) FindByUserID(ctx context.Context, userID uint) ([]*domain.Task, error) {

	return nil, nil
}
