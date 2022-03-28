package rest

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/laghodessa/togo/app"
)

func RegisterTasks(api fiber.Router, todoUC *app.TodoUsecase) {
	h := &taskHandlers{}

	tasks := api.Group("/tasks")
	tasks.Post("/", h.addTask(todoUC))
}

type taskHandlers struct{}

func (*taskHandlers) addTask(todoUC *app.TodoUsecase) fiber.Handler {
	type Request struct {
		Task struct {
			UserID  string `json:"userId"`
			Message string `json:"message"`
		}
		TimeZone string `json:"timeZone"`
	}

	return func(c *fiber.Ctx) error {
		var req Request
		if err := c.BodyParser(&req); err != nil {
			return ErrMalformedRequestBody
		}

		task, err := todoUC.AddTask(c.Context(), app.AddTask{
			Task: app.Task{
				UserID:  req.Task.UserID,
				Message: req.Task.Message,
			},
			TimeZone: req.TimeZone,
		})
		if err != nil {
			return fmt.Errorf("add task: %w", err)
		}
		return c.Status(http.StatusCreated).JSON(task)
	}
}
