package http

import (
	"net/http"
	"togo/internal/domain"

	"github.com/labstack/echo/v4"
)

func (s *httpServer) Register(c echo.Context) (err error) {
	u := new(domain.User)
	if err = c.Bind(u); err != nil {
		return err
	}
	s.userService.CreateUser(c.Request().Context(), &domain.User{
		FullName: u.FullName,
		Username: u.Username,
		Password: u.Password,
	})
	return nil
}

func (s *httpServer) Login(c echo.Context) (err error) {
	u := new(domain.LoginCredential)
	if err = c.Bind(u); err != nil {
		return
	}
	result, err := s.authService.Login(c.Request().Context(), u)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}
