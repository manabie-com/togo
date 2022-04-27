package repositories

import (
	"manabie.com/internal/models"
	"context"
)

type UserRepositoryI interface {
	/// return common.NotFound if iId not found
	FetchUserById(iContext context.Context, iId int) (models.User, error)
}
