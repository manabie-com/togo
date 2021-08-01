package router

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"togo/config"
)

type router struct {
	*config.Config
	*sql.DB
}

func NewRouter(config *config.Config, DB *sql.DB) *router {
	return &router{Config: config, DB: DB}
}

func ping() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON(&fiber.Map{
			"msg": "pong",
		})
	}
}

func (r *router) Start() error {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/", ping())
	app.Get("/ping", ping())

	addr := fmt.Sprintf(":%d", r.Port)
	err := app.Listen(addr)
	if err != nil {
		return err
	}

	return nil
}
