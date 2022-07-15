package task

import (
	"github.com/datshiro/togo-manabie/internal/interfaces/domain"
	"github.com/datshiro/togo-manabie/internal/interfaces/task/delivery/http"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(e *echo.Echo, apiPrefix string, taskUC domain.TaskUseCase, userUC domain.UserUseCase) {
	handler := &http.TaskHandler{
		TUseCase:       taskUC,
		UUseCase:       userUC,
		NewPostRequest: http.NewPostRequest,
	}
	e.POST(apiPrefix, handler.HandlePost)
}
