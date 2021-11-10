package http

import (
	"github.com/labstack/echo/v4"
	"github.com/manabie-com/togo/internal/domain"
	"github.com/manabie-com/togo/internal/transport/http/request"
	"github.com/manabie-com/togo/internal/transport/http/response"
	"net/http"
)

// AuthHandler represent auth http handler
type AuthHandler struct {
	AuthUseCase domain.AuthUseCase
}

// NewAuthHandler initialize new auth endpoint
func NewAuthHandler(e *echo.Echo, authUseCase domain.AuthUseCase) {
	authHandler := &AuthHandler{
		AuthUseCase: authUseCase,
	}
	e.GET("/login", authHandler.Login)
}

func (t *AuthHandler) Login(ctx echo.Context) error {
	var userLoginRequest request.UserLoginRequest
	if err := (&echo.DefaultBinder{}).BindQueryParams(ctx, &userLoginRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
	}
	if userLoginRequest.Username == "" {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "username must be not null or empty",
		})
	}
	if userLoginRequest.Password == "" {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "password must be not null or empty",
		})
	}
	accessToken, err := t.AuthUseCase.SignIn(ctx.Request().Context(), userLoginRequest.Username, userLoginRequest.Password)
	if err != nil {
		return ctx.JSON(response.GetStatusCode(err), response.ErrorResponse{
			Message: err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"data": accessToken,
	})
}
