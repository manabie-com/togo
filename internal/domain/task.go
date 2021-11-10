package domain

import (
	"context"
	"github.com/manabie-com/togo/internal/domain/entity"
)

// TaskUseCase task use case
type TaskUseCase interface {
	// CreateTask create new task
	CreateTask(ctx context.Context, content string, username string) error
	GetTask(ctx context.Context, username string, date string) ([]entity.Task, error)
}

// TaskRepository task repository interface
type TaskRepository interface {
	Create(ctx context.Context, content string, username string) error
	CountTaskInDayByUsername(ctx context.Context, username string) (int, error)
	GetTaskByUsernameAndDate(ctx context.Context, username string, date string) ([]entity.Task, error)
}
