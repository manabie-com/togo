package services

import (
	"context"
	"fmt"
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
	user, err := s.userRepo.FindOne(ctx, &domain.User{ID: task.UserID})
	if err != nil {
		return nil, fmt.Errorf("taskService:Create: %w", err)
	}
	if _, err = s.taskLimitRepo.Increase(ctx, user.ID, user.TasksPerDay); err != nil {
		return nil, fmt.Errorf("taskService:Create: %w", err)
	}
	return s.taskRepo.Create(ctx, task)
}
func (s taskService) Update(ctx context.Context, filter, update *domain.Task) (*domain.Task, error) {
	task, err := s.taskRepo.Update(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("taskService:UpdateByID: %w", err)
	}
	return task, nil
}
func (s taskService) FindByUserID(ctx context.Context, userID uint) ([]*domain.Task, error) {
	tasks, err := s.taskRepo.Find(ctx, &domain.Task{UserID: userID})
	if err != nil {
		return nil, fmt.Errorf("taskService:FindByUserID: %w", err)
	}
	return tasks, nil
}
