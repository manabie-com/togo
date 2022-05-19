package repositories

import (
	"manabie.com/internal/models"
	"manabie.com/internal/common"
	"context"
)

type TaskRepositoryI interface {
	/// return common.NotFound if iUser not found
	CreateTaskForUser(iContext context.Context, iUser models.User, iTasks []models.Task) ([]models.Task, error)
	/// return common.NotFound if iUser not found
	FetchNumberOfTaskForUser(iContext context.Context, iUser models.User) (int, error)
	/// return common.NotFound if iUser not found
	FetchNumberOfTaskForUserCreatedOnDay(iContext context.Context, iUser models.User, iCreatedTime common.Time) (int, error)
}