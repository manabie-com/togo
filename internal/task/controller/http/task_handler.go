package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"togo/domain"
)

type TaskHandler struct {
	service domain.ITaskService
}

func NewTaskHandler(e *echo.Echo, service domain.ITaskService) {
	handler := &TaskHandler{
		service: service,
	}
	e.POST("/tasks", handler.Create)
}

func (h *TaskHandler) Create(c echo.Context) error {
	var taskParams domain.TaskParams

	if err := c.Bind(&taskParams); err != nil {
		return err
	}

	if err := c.Validate(&taskParams); err != nil {
		return err
	}

	task, err := h.service.Create(taskParams)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, domain.ResponseSuccess{Data: task})
}
