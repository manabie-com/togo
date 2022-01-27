package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
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
	app         *fiber.App
}

func NewServer(config *configs.Config, userHandler *v1.UserHandler, taskHandler *v1.TaskHandler) *Server {
	app := fiber.New(fiber.Config{
		IdleTimeout: config.IdleTimeout,
	})
	return &Server{
		cfg:         config,
		userHandler: userHandler,
		taskHandler: taskHandler,
		app:         app,
	}
}

func (s *Server) Run() {
	s.app.Use(
		logger.New(logger.Config{
			Format: "[${ip}]:${port} ${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
		}),
		recover.New(),
	)
	v1.MapRoutes(s.app, s.cfg, s.userHandler, s.taskHandler)

	// Listen from a different goroutine
	go func() {
		if err := s.app.Listen(fmt.Sprintf(":%d", s.cfg.Port)); err != nil {
			panic(err)
		}
	}()

	c := make(chan os.Signal, 1)   // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, // When an interrupt or termination signal is sent, notify the channel
		syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)

	_ = <-c // This blocks the main thread until an interrupt is received
	logrus.Info("Gracefully shutting down...")
	_ = s.app.Shutdown()

	// Your cleanup tasks go here
	logrus.Info("Fiber was successful shutdown.")
}
