package rest

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/laghodessa/togo/app"
)

func RegisterUsers(api fiber.Router, todoUC *app.TodoUsecase) {
	h := &usersHandler{}

	tasks := api.Group("/users")
	tasks.Post("/", h.addUser(todoUC))
}

type usersHandler struct{}

func (*usersHandler) addUser(todoUC *app.TodoUsecase) fiber.Handler {
	type Request struct {
		TaskDailyLimit int `json:"taskDailyLimit"`
	}

	return func(c *fiber.Ctx) error {
		var req Request
		if err := c.BodyParser(&req); err != nil {
			return ErrMalformedRequestBody
		}

		user, err := todoUC.AddUser(c.Context(), app.User{
			TaskDailyLimit: req.TaskDailyLimit,
		})
		if err != nil {
			return fmt.Errorf("add user: %w", err)
		}
		return c.Status(http.StatusCreated).JSON(User{
			ID:             user.ID,
			TaskDailyLimit: user.TaskDailyLimit,
		})
	}
}
