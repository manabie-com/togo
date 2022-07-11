package user

import (
	"path"

	"github.com/datshiro/togo-manabie/internal/interfaces/domain"
	"github.com/datshiro/togo-manabie/internal/interfaces/user/delivery/http"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(e *echo.Echo, apiPrefix string, uc domain.UserUseCase) {
	handler := &http.UserHandler{
		UUseCase:       uc,
		NewPostRequest: http.NewPostRequest,
		NewGetRequest:  http.NewGetRequest,
	}
	e.POST(apiPrefix, handler.HandlePost)
	e.GET(path.Join(apiPrefix, ":id"), handler.HandleGet)
}
