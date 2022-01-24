package repository

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trinhdaiphuc/togo/internal/entities"
	"github.com/trinhdaiphuc/togo/internal/infrastructure"
)

type TaskRepository interface {
	Create(ctx *fiber.Ctx, task *entities.Task) (*entities.Task, error)
}

type taskRepositoryImpl struct {
	db infrastructure.DB
}

func NewTaskRepository(db infrastructure.DB) TaskRepository {
	return &taskRepositoryImpl{
		db: db,
	}
}

func (t taskRepositoryImpl) Create(ctx *fiber.Ctx, task *entities.Task) (*entities.Task, error) {
	panic("implement me")
}
