package storages

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/model"
)

type Store interface {
	User() User
	Task() Task
	DropAllRecords()
}

type User interface {
	Get(ctx context.Context, userID string) (*model.User, error)
	Create(ctx context.Context, u *model.User) (*model.User, error)
}

type Task interface {
	RetrieveTasks(ctx context.Context, userID string, createdDate sql.NullString) ([]*model.Task, error)
	AddTask(ctx context.Context, userID string, t *model.Task) (*model.Task, error)
	CountTasksByUser(ctx context.Context, userID string) (int, error)
}
