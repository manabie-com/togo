package server

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/trangmaiq/togo/internal/config"
	"github.com/trangmaiq/togo/internal/registry"
	"github.com/trangmaiq/togo/internal/server/handler"
)

func Start(cfg *config.ToGo) {
	h := handler.New(registry.Registry())

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/tasks", h.CreateTasks())

	err := e.Start(fmt.Sprintf("0.0.0.0:%d", cfg.ServicePort))
	if err != nil {
		logrus.WithField("err", err).Fatal("start http server failed")
	}
}
