package tasks

import (
	"net/http"
	"strconv"

	"manabie/todo/models"
	"manabie/todo/pkg/utils"
	"manabie/todo/service/task"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type StatusResponse struct {
	Status string `json:"status"`
}

type TaskIndexResponse struct {
	Tasks []*models.Task `json:"tasks"`
}

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
		return utils.ResponseFailure(c, http.StatusBadRequest, err)
	}

	tasks, err := h.Task.Index(c.Request().Context(), memberID, date)
	if err != nil {
		return utils.ResponseFailure(c, http.StatusInternalServerError, err)
	}

	res := TaskIndexResponse{
		Tasks: tasks,
	}

	return utils.ResponseSuccess(c, res)
}

func (h *handler) Show(c echo.Context) error {
	id := c.Param("id")

	taskID, err := strconv.Atoi(id)
	if err != nil {
		return utils.ResponseFailure(c, http.StatusBadRequest, err)
	}

	tk, err := h.Task.Show(c.Request().Context(), taskID)
	if err != nil {
		return utils.ResponseFailure(c, http.StatusNotFound, err)
	}

	return utils.ResponseSuccess(c, tk)
}

func (h *handler) Create(c echo.Context) error {

	id := c.Param("id")
	tk := new(models.Task)

	if err := c.Bind(tk); err != nil {
		return utils.ResponseFailure(c, http.StatusBadRequest, err)
	}

	// TODO Add validate
	if no, _ := strconv.Atoi(id); no != tk.MemberID {
		return utils.ResponseFailure(c, http.StatusBadRequest, errors.New("Member Id incorrect"))
	}

	if err := h.Task.Create(c.Request().Context(), tk); err != nil {
		return utils.ResponseFailure(c, http.StatusInternalServerError, err)
	}

	return utils.ResponseSuccess(c, StatusResponse{
		Status: "ok",
	})
}

func (h *handler) Update(c echo.Context) error {

	tk := new(models.Task)

	if err := c.Bind(tk); err != nil {
		return utils.ResponseFailure(c, http.StatusBadRequest, err)
	}

	if err := h.Task.Update(c.Request().Context(), tk); err != nil {
		return utils.ResponseFailure(c, http.StatusInternalServerError, err)
	}

	return utils.ResponseSuccess(c, StatusResponse{
		Status: "ok",
	})
}

func (h *handler) Delete(c echo.Context) error { return nil }
