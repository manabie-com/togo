package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/trinhdaiphuc/logger"
	"github.com/trinhdaiphuc/togo/configs"
	v1 "github.com/trinhdaiphuc/togo/internal/api/v1"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	cfg         *configs.Config
	userHandler *v1.UserHandler
	taskHandler *v1.TaskHandler
}

func NewServer(config *configs.Config, userHandler *v1.UserHandler, taskHandler *v1.TaskHandler) *Server {
	return &Server{
		cfg:         config,
		userHandler: userHandler,
		taskHandler: taskHandler,
	}
}

func (s *Server) Run() {
	app := fiber.New(fiber.Config{
		IdleTimeout: s.cfg.IdleTimeout,
	})
	app.Use(logger.FiberMiddleware())
	v1.MapRoutes(app, s.userHandler, s.taskHandler)

	// Listen from a different goroutine
	go func() {
		if err := app.Listen(fmt.Sprintf(":%d", s.cfg.Port)); err != nil {
			panic(err)
		}
	}()

	c := make(chan os.Signal, 1)   // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, // When an interrupt or termination signal is sent, notify the channel
		syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)

	_ = <-c // This blocks the main thread until an interrupt is received
	logrus.Info("Gracefully shutting down...")
	_ = app.Shutdown()

	// Your cleanup tasks go here
	logrus.Info("Fiber was successful shutdown.")
}
