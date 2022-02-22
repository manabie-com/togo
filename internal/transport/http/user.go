package http

import (
	"net/http"
	"togo/internal/domain"

	"github.com/labstack/echo/v4"
)

func (s *httpServer) GetMe(c echo.Context) error {
	currentUser, err := s.getCurrentUser(c)
	if err != nil {
		return err
	}
	user, err := s.userService.GetUserByID(c.Request().Context(), currentUser.ID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

type updateUserDTO struct {
	FullName    string `json:"fullName"`
	TasksPerDay int    `json:"tasksPerDay"`
}

func (s *httpServer) UpdateMe(c echo.Context) error {
	u := new(updateUserDTO)
	if err := c.Bind(u); err != nil {
		return err
	}
	currentUser, err := s.getCurrentUser(c)
	if err != nil {
		return err
	}
	user, err := s.userService.UpdateByID(c.Request().Context(), currentUser.ID, &domain.User{
		FullName:    u.FullName,
		TasksPerDay: u.TasksPerDay,
	})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}
