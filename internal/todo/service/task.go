package service

import (
	"context"
	"time"

	d "github.com/manabie-com/togo/internal/todo/domain"
)

type TaskService struct {
	TaskRepo d.TaskRepository
}

func NewTaskService(taskRepo d.TaskRepository) *TaskService {
	return &TaskService{taskRepo}
}

func (s *TaskService) ListTaskForUser(ctx context.Context, userID int, dateStr string) ([]*d.Task, error) {
	if _, err := time.Parse("2006-01-02", dateStr); err != nil {
		dateStr = time.Now().Format("2006-01-02")
	}

	return s.TaskRepo.GetTasksForUser(ctx, userID, dateStr)
}

func (s *TaskService) CreateTaskForUser(ctx context.Context, userID int, param d.TaskCreateParam) (*d.Task, error) {
	task, err := s.TaskRepo.CreateTaskForUser(ctx, userID, param)
	if task == nil && err == nil {
		return nil, d.ErrTaskLimitReached
	}

	return task, err
}
