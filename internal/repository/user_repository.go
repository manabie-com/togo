package repository

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trinhdaiphuc/togo/internal/entities"
	"github.com/trinhdaiphuc/togo/internal/infrastructure"
)

type UserRepository interface {
	GetUser(ctx *fiber.Ctx, userName string) (*entities.User, error)
}

type userRepositoryImpl struct {
	db infrastructure.DB
}

func NewUserRepository(db infrastructure.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (u userRepositoryImpl) GetUser(ctx *fiber.Ctx, userName string) (*entities.User, error) {
	panic("implement me")
}
