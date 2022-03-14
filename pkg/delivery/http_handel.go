package delivery

import (
	"context"
	"github.com/labstack/echo/v4"
	"togo.com/pkg/model"
	"togo.com/pkg/usecase"
)

type handle struct {
	taskUc      usecase.TaskUseCase
	authorizeUc usecase.AuthorizeUseCase
}

func HttpHandel(e *echo.Echo, taskUc usecase.TaskUseCase, authorizeUc usecase.AuthorizeUseCase) {
	h := &handle{
		taskUc:      taskUc,
		authorizeUc: authorizeUc,
	}
	e.POST("/login", h.Login)
	e.POST("/task", h.AddTask)
}

func (h handle) Login(c echo.Context) error {
	req := model.LoginRequest{}
	err := c.Bind(&req)
	if err != nil {
		return model.ResponseWithError(c, err)
	}
	resp, err := h.authorizeUc.Login(context.Background(), req)
	return model.ResponseSuccess(c, resp)
}

func (h handle) AddTask(c echo.Context) error {
	req := model.AddTaskRequest{}
	err := c.Bind(&req)
	token := c.Request().Header.Get("Authorization")
	userId, errString := usecase.ValidateToken(token)
	if errString != "" {
		return model.ResponseWithError(c, errString)
	}
	if err != nil {
		return model.ResponseWithError(c, err)
	}
	resp, err := h.taskUc.AddTask(context.Background(), userId, req)
	if err != nil {
		return model.ResponseWithError(c, err)
	}
	return model.ResponseSuccess(c, resp)
}
