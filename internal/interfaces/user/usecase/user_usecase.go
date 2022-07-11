package usecase

import (
	"context"
	"database/sql"

	"github.com/datshiro/togo-manabie/internal/interfaces/domain"
	"github.com/datshiro/togo-manabie/internal/interfaces/models"
	postgres_user "github.com/datshiro/togo-manabie/internal/interfaces/user/repository/postgres"
)

func NewUserUseCase(dbc *sql.DB) domain.UserUseCase {
	return &userUseCase{
		DB:       dbc,
		userRepo: postgres_user.NewUserRepository(),
	}
}

type userUseCase struct {
	userRepo domain.UserRepository
	DB       *sql.DB
}

func (u *userUseCase) GetUser(ctx context.Context, userID int) (*models.User, error) {
	return u.userRepo.GetUser(ctx, u.DB, userID)
}

func (u *userUseCase) AddTask(ctx context.Context, user *models.User, task *models.Task) error {
	return u.userRepo.AddTask(ctx, u.DB, user, task)
}

func (t *userUseCase) CreateUser(ctx context.Context, m *models.User) error {
	return t.userRepo.CreateOne(ctx, t.DB, m)
}
