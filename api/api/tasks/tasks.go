package tasks

import (
	"manabie/todo/service/task"

	"github.com/labstack/echo/v4"
)

type handler struct {
	task task.TaskService
}

func NewTaskHandler(e *echo.Echo, ts task.TaskService) {
	h := &handler{
		task: ts,
	}
	e.GET("/tasks", h.Index)
	e.POST("/tasks", h.Create)
	e.GET("/tasks/:id", h.Show)
	e.PUT("/tasks/:id", h.Update)
	e.DELETE("/tasks/:id", h.Delete)
}

func (h *handler) Index(c echo.Context) error  { return nil }
func (h *handler) Show(c echo.Context) error   { return nil }
func (h *handler) Create(c echo.Context) error { return nil }
func (h *handler) Update(c echo.Context) error { return nil }
func (h *handler) Delete(c echo.Context) error { return nil }
