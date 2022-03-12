package delivery

import (
	"context"
	"github.com/labstack/echo/v4"
	"log"
	"togo.com/pkg/model"
	"togo.com/pkg/usecase"
)

type handle struct {
	taskUc usecase.TaskUseCase
	userUc usecase.UserUseCase
}

func HttpHandel(e *echo.Echo, taskUc usecase.TaskUseCase, userUc usecase.UserUseCase) {
	h := &handle{
		taskUc: taskUc,
		userUc: userUc,
	}
	e.POST("/login", h.Login)
}

func (h handle) Login(c echo.Context) error {
	req := model.LoginRequest{}
	err := c.Bind(&req)
	if err != nil {
		log.Fatal(err)
		return model.ResponseWithError(c, err)
	}
	res := h.userUc.Login(context.Background(), req)
	return model.ResponseSuccess(c, res)
}
