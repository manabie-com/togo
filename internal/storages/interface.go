package storages

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/entities"
)

type Task interface {
	RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*entities.Task, error)
	AddTask(ctx context.Context, t *entities.Task) error
	ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool
}
