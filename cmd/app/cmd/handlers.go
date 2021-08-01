package cmd

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"togo/config"
)

type ServerConfig struct {
	*config.Config
	*sql.DB
}
type server struct {
	*ServerConfig
}

func NewServer(serverConfig *ServerConfig) *server {
	return &server{ServerConfig: serverConfig}
}

func ping() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON(&fiber.Map{
			"msg": "pong",
		})
	}
}

func (s *server) Start() error {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/", ping())
	app.Get("/ping", ping())

	addr := fmt.Sprintf(":%d", s.Port)
	err := app.Listen(addr)
	if err != nil {
		return err
	}

	return nil
}
