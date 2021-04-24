package http

import (
	"github.com/labstack/echo/v4"
	"github.com/manabie-com/togo/internal/domain"
)

type HttpAPI struct {
	taskUseCase domain.TaskUseCase
	userUseCase domain.AuthUseCase
}

func NewTaskHttpServer(e echo.Echo) HttpAPI {
	return HttpAPI{}
}

func (h HttpAPI) Login(ctx echo.Context) error {
	return nil

}

func (h HttpAPI) AddTask(ctx echo.Context) error {
	return nil

}

func (h HttpAPI) ListTasks(ctx echo.Context) error {
	return nil

}
