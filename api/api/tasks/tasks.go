package tasks

import (
	"strconv"

	"manabie/todo/models"
	"manabie/todo/pkg/apiutils"
	"manabie/todo/service/task"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type handler struct {
	Task task.TaskService
}

func NewTaskHandler(e *echo.Echo, ts task.TaskService) {
	h := &handler{
		Task: ts,
	}
	e.GET("/users/:id/tasks", h.Index)
	e.POST("/users/:id/tasks", h.Create)

	e.GET("/tasks/:id", h.Show)
	e.PUT("/tasks/:id", h.Update)
	e.DELETE("/tasks/:id", h.Delete)
}

func (h *handler) Index(c echo.Context) error {
	id := c.Param("id")
	date := c.QueryParam("date")

	memberID, err := strconv.Atoi(id)
	if err != nil {
		return apiutils.ResponseFailure(c, errors.Wrap(apiutils.ErrInvalidValue, err.Error()))
	}

	tasks, err := h.Task.Index(c.Request().Context(), memberID, date)
	if err != nil {
		return apiutils.ResponseFailure(c, err)
	}

	return apiutils.ResponseSuccess(c, models.TaskIndexResponse{
		Tasks: tasks,
	})
}

func (h *handler) Show(c echo.Context) error {
	id := c.Param("id")

	taskID, err := strconv.Atoi(id)
	if err != nil {
		return apiutils.ResponseFailure(c, errors.Wrap(apiutils.ErrInvalidValue, err.Error()))
	}

	tk, err := h.Task.Show(c.Request().Context(), taskID)
	if err != nil {
		return apiutils.ResponseFailure(c, err)
	}

	return apiutils.ResponseSuccess(c, tk)
}

func (h *handler) Create(c echo.Context) error {

	id := c.Param("id")
	tk := new(models.TaskCreateRequest)

	memberID, err := strconv.Atoi(id)
	if err != nil {
		return apiutils.ResponseFailure(c, errors.Wrap(apiutils.ErrInvalidValue, err.Error()))
	}

	if err := c.Bind(tk); err != nil {
		return apiutils.ResponseFailure(c, errors.Wrap(apiutils.ErrInvalidValue, err.Error()))
	}

	if err := h.Task.Create(c.Request().Context(), memberID, tk); err != nil {
		return apiutils.ResponseFailure(c, err)
	}

	return apiutils.ResponseSuccess(c, models.StatusResponse{
		Status: "ok",
	})
}

func (h *handler) Update(c echo.Context) error {
	id := c.Param("id")
	tk := new(models.Task)

	ID, err := strconv.Atoi(id)
	if err != nil {
		return apiutils.ResponseFailure(c, errors.Wrap(apiutils.ErrInvalidValue, err.Error()))
	}

	if err := c.Bind(tk); err != nil {
		return apiutils.ResponseFailure(c, errors.Wrap(apiutils.ErrInvalidValue, err.Error()))
	}

	if err := h.Task.Update(c.Request().Context(), ID, tk); err != nil {
		return apiutils.ResponseFailure(c, err)
	}

	return apiutils.ResponseSuccess(c, models.StatusResponse{
		Status: "ok",
	})
}

func (h *handler) Delete(c echo.Context) error {
	id := c.Param("id")

	taskID, err := strconv.Atoi(id)
	if err != nil {
		return apiutils.ResponseFailure(c, errors.Wrap(apiutils.ErrInvalidValue, err.Error()))
	}

	if err := h.Task.Delete(c.Request().Context(), taskID); err != nil {
		return apiutils.ResponseFailure(c, err)
	}

	return apiutils.ResponseSuccess(c, models.StatusResponse{
		Status: "ok",
	})
}
