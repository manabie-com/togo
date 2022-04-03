package auth

import (
	"net/http"
	"strings"

	"github.com/TrinhTrungDung/togo/internal/model"
	"github.com/labstack/echo/v4"
)

// Service represents auth application interface
type Service interface {
	Register(RegisterData) (*model.User, error)
	LoginUsername(LoginUsernameData) (*model.AuthToken, error)
	LoginEmail(LoginEmailData) (*model.AuthToken, error)
}

// HTTP represents auth http service
type HTTP struct {
	svc Service
}

// RegisterData contains user's registration data from JSON request
type RegisterData struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

// LoginUsernameData contains user's login data using username from JSON request
type LoginUsernameData struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginEmailData contains user's login data using email from JSON request
type LoginEmailData struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func NewHTTP(svc Service, eg *echo.Group) {
	h := HTTP{svc}

	eg.POST("/register", h.register)
	eg.POST("/login-username", h.loginUsername)
	eg.POST("/login-email", h.loginEmail)
}

func (h *HTTP) register(c echo.Context) error {
	body := RegisterData{}
	if err := c.Bind(&body); err != nil {
		return err
	}
	body.Email = strings.TrimSpace(body.Email)
	body.FirstName = strings.TrimSpace(body.FirstName)
	body.LastName = strings.TrimSpace(body.LastName)

	resp, err := h.svc.Register(body)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HTTP) loginUsername(c echo.Context) error {
	body := LoginUsernameData{}
	if err := c.Bind(&body); err != nil {
		return err
	}
	body.Username = strings.TrimSpace(body.Username)

	resp, err := h.svc.LoginUsername(body)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HTTP) loginEmail(c echo.Context) error {
	body := LoginEmailData{}
	if err := c.Bind(&body); err != nil {
		return err
	}
	body.Email = strings.TrimSpace(body.Email)

	resp, err := h.svc.LoginEmail(body)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}
