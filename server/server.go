package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"togo/configs"
	"togo/internal/middleware"
	"togo/internal/response"
	taskHandler "togo/internal/task/handler"
	taskRepo "togo/internal/task/repository"
	taskService "togo/internal/task/service"
	userHandler "togo/internal/user/handler"
	userRepo "togo/internal/user/repository"
	userservice "togo/internal/user/service"
	"togo/internal/validator"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Server struct {
	Config *configs.Config
	Logger *zap.Logger
	Db     *gorm.DB
}

func (s *Server) Start() {
	e := echo.New()
	e.Validator = validator.NewValidator()
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		Skipper:      echoMiddleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut},
		MaxAge:       86400,
	}))

	userRepo := userRepo.NewUserRepository(s.Db)
	taskRepo := taskRepo.NewTaskRepository(s.Db)

	userService := userservice.NewUserService(userRepo)
	taskService := taskService.NewTaskService(taskRepo, userRepo)

	e.GET("/health", func(c echo.Context) error {
		return response.Success(c, map[string]interface{}{
			"time":       time.Now(),
			"ip_address": c.RealIP(),
		})
	})

	userGroup := e.Group("/users", middleware.Middleware())
	userHandler.NewUserHandler(userGroup, userService)

	taskGroup := userGroup.Group("/:id/tasks", middleware.TaskMiddlerware())
	taskHandler.NewTaskHandler(taskGroup, taskService)

	go func() {
		if err := e.Start(":" + s.Config.Server.Port); err != nil && err != http.ErrServerClosed {
			s.Logger.Sugar().Fatalf("[CRITICAL] Shutting down the server: %+v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		s.Logger.Sugar().Fatalf("[CRITICAL] Server shutdown failed: %+v", err)
	}
}
