package storages

import (
	"context"

	"github.com/manabie-com/togo/internal/entity"
)

type TaskStorage interface {
	RetrieveTasks(ctx context.Context, userID, createdDate string) ([]*entity.Task, error)
	AddTask(ctx context.Context, task *entity.Task) error
	GetNumberOfTasks(ctx context.Context, userID, date string) (int, error)
}

type UserStorage interface {
	FindByID(ctx context.Context, userID string) (*entity.User, error)
}
