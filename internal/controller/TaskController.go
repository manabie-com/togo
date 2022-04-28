package controller

import (
	"togo/internal/connect"
	"togo/internal/services"

	"github.com/gofiber/fiber/v2"
)

func CreateTask(ctx *fiber.Ctx) error {
	db := connect.DB

	c := services.Con{
		Db:  db,
		Ctx: ctx,
	}

	return c.CreateTask()
}
