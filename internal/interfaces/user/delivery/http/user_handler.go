package http

import (
	"net/http"

	"github.com/datshiro/togo-manabie/internal/infras/errors"
	"github.com/datshiro/togo-manabie/internal/interfaces/domain"
	"github.com/datshiro/togo-manabie/internal/interfaces/models"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
)

type UserHandler struct {
	UUseCase       domain.UserUseCase
	NewPostRequest func(c echo.Context) PostRequest
	NewGetRequest  func(c echo.Context) GetRequest
}

// Create Handler
func (h UserHandler) HandlePost(c echo.Context) error {
	request := h.NewPostRequest(c)
	if err := request.Bind(); err != nil {
		return errors.InvalidRequest
	}
	if err := request.Validate(); err != nil {
		return err
	}

	user := &models.User{
		Name:  request.GetName(),
		Email: null.StringFrom(request.GetEmail()),
		Quota: request.GetQuota(),
	}
	ctx := c.Request().Context()

	if err := h.UUseCase.CreateUser(ctx, user); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

// Get Handler
func (h UserHandler) HandleGet(c echo.Context) error {
	request := h.NewGetRequest(c)
	if err := request.Bind(); err != nil {
		return errors.InvalidRequest
	}

	user, err := h.UUseCase.GetUser(c.Request().Context(), request.GetID())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, NewUserResponse(user))
}
