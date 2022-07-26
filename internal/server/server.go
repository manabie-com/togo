package server

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/trangmaiq/togo/internal/config"
	"github.com/trangmaiq/togo/internal/registry"
	"github.com/trangmaiq/togo/internal/server/handler"
	"github.com/trangmaiq/togo/pkg/graceful"
)

func StartWithGracefulShutdown(cfg *config.ToGo) {
	e := echo.New()
	err := graceful.Graceful(func() error {
		h := handler.New(registry.Registry())

		e.Use(middleware.Logger())
		e.Use(middleware.Recover())

		e.POST("/tasks", h.CreateTasks())

		err := e.Start(fmt.Sprintf("0.0.0.0:%d", cfg.ServicePort))
		if err != nil {
			return fmt.Errorf("start http server failed: %w", err)
		}

		return nil
	}, func(ctx context.Context) error {
		// TODO: Handle errors
		_ = e.Close()
		_ = registry.Close()

		return nil
	})
	if err != nil {
		logrus.WithField("err", err).Fatal("gracefully shutdown")
	}
}
