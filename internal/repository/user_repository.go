package repository

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/trinhdaiphuc/togo/database/ent"
	"github.com/trinhdaiphuc/togo/database/ent/user"
	"github.com/trinhdaiphuc/togo/internal/dto"
	"github.com/trinhdaiphuc/togo/internal/entities"
	"github.com/trinhdaiphuc/togo/internal/infrastructure"
	"github.com/trinhdaiphuc/togo/pkg/helper"
)

type UserRepository interface {
	GetUserByName(ctx context.Context, username string) (*entities.User, error)
	CreateUser(ctx context.Context, user *entities.User) (*entities.User, error)
}

type userRepositoryImpl struct {
	db *infrastructure.DB
}

var ErrUserNotFound = fiber.NewError(fiber.StatusNotFound, "User not found")

func NewUserRepository(db *infrastructure.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (u *userRepositoryImpl) GetUserByName(ctx context.Context, username string) (*entities.User, error) {
	resp, err := u.db.User.Query().Where(user.UsernameEQ(username)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return dto.User2UserEntity(resp), nil
}

func (u *userRepositoryImpl) CreateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	password, err := helper.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	resp, err := u.db.User.
		Create().
		SetUsername(user.Username).
		SetPassword(password).
		SetTaskLimit(user.TaskLimit).
		Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			return nil, fiber.NewError(fiber.StatusConflict, "Username already exists")
		}
		return nil, err
	}
	return dto.User2UserEntity(resp), nil
}
