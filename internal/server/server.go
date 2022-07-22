package server

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/trangmaiq/togo/internal/config"
)

func Start(cfg *config.ToGo) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/tasks", create)

	err := e.Start(fmt.Sprintf("0.0.0.0:%d", cfg.ServicePort))
	if err != nil {
		logrus.WithField("err", err).Fatal("start http server failed")
	}
}

func create(c echo.Context) error {
	err := c.JSON(200, map[string]string{
		"status": "succeeded",
	})
	if err != nil {
		logrus.WithField("err", err).Error("send JSON response failed")
		return fmt.Errorf("send JSON response failed: %w", err)
	}

	return nil
}
