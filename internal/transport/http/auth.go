package http

import (
	"net/http"
	"strings"
	"togo/internal/domain"

	"github.com/labstack/echo/v4"
)

const (
	currentUserKey2 = "CURRENT_USER"
)

func (s *httpServer) authGuard(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract auth token from request
		authHeader := c.Request().Header.Get("Authorization")
		// Validate and verify token
		authHeader = strings.Trim(authHeader, " ")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return echo.NewHTTPError(http.StatusUnauthorized, domain.ErrUnauthorized.Error())
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		result, err := s.authService.VerifyToken(c.Request().Context(), token)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, domain.ErrUnauthorized.Error())
		}
		// Attach auth user
		c.Set(currentUserKey2, result.Payload)
		return next(c)
	}
}

func (s *httpServer) getCurrentUser(c echo.Context) (*domain.User, error) {
	user, ok := c.Get(currentUserKey2).(*domain.User)
	if !ok {
		return nil, domain.ErrUnauthorized
	}
	return user, nil
}

type registerDTO struct {
	FullName    string `json:"fullName,required"`
	Username    string `json:"username,required"`
	Password    string `json:"password,required"`
	TasksPerDay int    `json:"tasksPerDay,required"`
}

func (s *httpServer) Register(c echo.Context) (err error) {
	u := new(registerDTO)
	if err = c.Bind(u); err != nil {
		return err
	}
	if _, err = s.userService.CreateUser(c.Request().Context(), &domain.User{
		FullName:    strings.Trim(u.FullName, " "),
		Username:    strings.Trim(u.Username, " "),
		Password:    u.Password,
		TasksPerDay: u.TasksPerDay,
	}); err != nil {
		return
	}
	return c.String(http.StatusCreated, "")
}

func (s *httpServer) Login(c echo.Context) (err error) {
	u := new(domain.LoginCredential)
	if err = c.Bind(u); err != nil {
		return
	}
	result, err := s.authService.Login(c.Request().Context(), u)
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, result)
}
