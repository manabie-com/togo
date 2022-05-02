package taskservice

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

import (
	"context"
	"todo/internal/entities"
	"todo/internal/repository"
)

type TaskService interface {
	CreateTask(ctx context.Context, task *entities.Task) (*entities.Task, error)
	GetTask(ctx context.Context, id int) (*entities.Task, error)
	GetTasks(ctx context.Context, filter *entities.TaskFilter) (*entities.Tasks, error)
	UpdateTask(ctx context.Context, task *entities.Task) (*entities.Task, error)
	DeleteTask(ctx context.Context, id int) error
}

type taskService struct {
	taskRepo repository.TaskRepository
}

func NewTaskService(taskRepo repository.TaskRepository) TaskService {
	return &taskService{
		taskRepo: taskRepo,
	}
}
