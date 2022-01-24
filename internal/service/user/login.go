package userservice

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trinhdaiphuc/togo/internal/entities"
)

func (u userService) Login(ctx *fiber.Ctx, user *entities.User) (*entities.User, error) {
	panic("implement me")
}
