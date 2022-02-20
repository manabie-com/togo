package http

import (
	"net/http"
	"togo/internal/domain"

	"github.com/labstack/echo/v4"
)

type getUserDTO struct {
	ID uint `param:"id"`
}

func (s *httpServer) GetUser(c echo.Context) error {
	u := new(getUserDTO)
	if err := c.Bind(u); err != nil {
		return err
	}
	user, err := s.userService.GetUserByID(c.Request().Context(), u.ID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

type updateUserDTO struct {
	ID          uint   `param:"id"`
	FullName    string `json:"fullName"`
	TasksPerDay int    `json:"tasksPerDay"`
}

func (s *httpServer) UpdateUser(c echo.Context) error {
	u := new(updateUserDTO)
	if err := c.Bind(u); err != nil {
		return err
	}
	user, err := s.userService.UpdateByID(c.Request().Context(), u.ID, &domain.User{
		FullName:    u.FullName,
		TasksPerDay: u.TasksPerDay,
	})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}
