package taskservice

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trinhdaiphuc/togo/internal/entities"
)

func (t taskService) CreateTask(ctx *fiber.Ctx, task *entities.Task) (*entities.Task, error) {
	panic("implement me")
}
