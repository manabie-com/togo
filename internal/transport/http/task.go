package http

import (
	"net/http"
	"togo/internal/domain"

	"github.com/labstack/echo/v4"
)

type taskCreateDTO struct {
	Content string `json:"content"`
}

func (s httpServer) AddTask(c echo.Context) error {
	u := new(taskCreateDTO)
	if err := c.Bind(u); err != nil {
		return err
	}
	currentUser, err := s.getCurrentUser(c)
	if err != nil {
		return err
	}
	task, err := s.taskService.Create(c.Request().Context(), &domain.Task{
		UserID:  currentUser.ID,
		Content: u.Content,
	})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, task)
}

func (s httpServer) GetTasks(c echo.Context) error {
	currentUser, err := s.getCurrentUser(c)
	if err != nil {
		return err
	}
	tasks, err := s.taskService.FindByUserID(c.Request().Context(), currentUser.ID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, tasks)
}

type taskUpdateDTO struct {
	ID      uint   `param:"id"`
	Content string `json:"content"`
}

func (s httpServer) UpdateTask(c echo.Context) error {
	u := new(taskUpdateDTO)
	if err := c.Bind(u); err != nil {
		return err
	}
	currentUser, err := s.getCurrentUser(c)
	if err != nil {
		return err
	}
	task, err := s.taskService.Update(c.Request().Context(),
		&domain.Task{
			ID:     u.ID,
			UserID: currentUser.ID,
		},
		&domain.Task{
			Content: u.Content,
		})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, task)
}
