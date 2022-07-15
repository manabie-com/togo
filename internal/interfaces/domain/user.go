package domain

import (
	"context"

	"github.com/datshiro/togo-manabie/internal/interfaces/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UserRepository interface {
	CreateOne(context.Context, boil.ContextExecutor, *models.User) error
	GetUser(context.Context, boil.ContextExecutor, int) (*models.User, error)
	AddTask(context.Context, boil.ContextExecutor, *models.User, *models.Task) error
}

type UserUseCase interface {
	CreateUser(context.Context, *models.User) error
	GetUser(context.Context, int) (*models.User, error)
	AddTask(context.Context, *models.User, *models.Task) error
}
