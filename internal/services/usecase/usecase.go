package usecase

import (
	"context"

	"github.com/manabie-com/togo/internal/app/models"
)

type Usecase interface {
	GetAuthToken(ctx context.Context, username string, pwd string) (string, error)
	RetrieveTasks(ctx context.Context, userID uint64, createdDate string) ([]*models.Task, error)
	AddTask(ctx context.Context, userID uint64, task *models.Task) (*models.Task, error)
}
