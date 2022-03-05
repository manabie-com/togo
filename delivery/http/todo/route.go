package todo

import (
	"github.com/labstack/echo/v4"

	"github.com/khangjig/togo/usecase"
)

type Route struct {
	useCase *usecase.UseCase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{useCase}

	group.POST("", r.create)
	group.GET("", r.getList)
	group.PUT("/:id", r.update)
	group.GET("/:id", r.getByID)
}
