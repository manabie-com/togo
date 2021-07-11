package storages

import (
	"context"

	"github.com/manabie-com/togo/internal/entity"
)

type TaskStorage interface {
	RetrieveTasks(ctx context.Context, userID, createdDate string) ([]*entity.Task, error)
	AddTask(ctx context.Context, task *entity.Task) error
}

type UserStorage interface {
	Find(ctx context.Context, userID, pw string) (*entity.User, error)
}
