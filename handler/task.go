package handler

import (
	"togo/service"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	taskService *service.TaskService
}

func NewTaskHandler(
	router *echo.Group,
	taskService *service.TaskService,
) {
	handler := &TaskHandler{taskService}
	router.POST("", handler.Create)
}

func (h *TaskHandler) Create(e echo.Context) error {
	return nil
}
