package storages

import (
	"context"

	"github.com/manabie-com/togo/internal/storages/entities"
)

type Store interface {
	RetrieveTasks(ctx context.Context, userID, createdDate string, limit, offset int) ([]*entities.Task, error)
	AddTask(ctx context.Context, t *entities.Task) error
	ValidateUser(ctx context.Context, userID, pwd string) bool
}
