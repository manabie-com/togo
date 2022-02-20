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
	currentUser := c.Get(currentUserKey).(*domain.User)
	task, err := s.taskService.Create(c.Request().Context(), &domain.Task{
		UserID:  currentUser.ID,
		Content: u.Content,
	})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, task)
}
