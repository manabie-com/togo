package storages

import (
	"context"

	"github.com/manabie-com/togo/internal/app/models"
)

type Storages interface {
	ValidateUser(ctx context.Context, username string, pwd string) (*models.User, error)
	RetrieveTasks(ctx context.Context, userID uint64, createdDate string) ([]*models.Task, error)
	AddTask(ctx context.Context, task *models.Task) error
}
