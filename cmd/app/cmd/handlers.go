package cmd

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	jwtware "github.com/gofiber/jwt/v2"
	"togo/config"
	"togo/internal/rest"
)

type server struct {
	SC *config.ServerConfig
}

func NewServer(serverConfig *config.ServerConfig) *server {
	return &server{SC: serverConfig}
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

	app.Post("/login", rest.UserLogin(s.SC))
	app.Post("/signup", rest.UserSignup(s.SC))

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
		SigningKey: []byte(s.SC.JwtSecret),
	}))

	tasks := app.Group("/tasks")
	tasks.Post("", rest.CreateTask(s.SC))
	tasks.Get("/:id", rest.GetTask(s.SC))
	tasks.Get("", rest.ListTask(s.SC))
	tasks.Patch("/:id", rest.UpdateTask(s.SC))
	tasks.Patch("/:id", rest.UpdateTask(s.SC))
	tasks.Delete("/:id", rest.DeleteTask(s.SC))

	addr := fmt.Sprintf(":%d", s.SC.Port)
	err := app.Listen(addr)
	if err != nil {
		return err
	}

	return nil
}
