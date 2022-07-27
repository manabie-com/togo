package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"togo/configs"
	"togo/internal/handler"
	"togo/internal/middleware"
	"togo/internal/repository"
	"togo/internal/response"
	"togo/internal/service"
	"togo/pkg/databases"
	"togo/pkg/logger"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func init() {
	logger.NewLogger()
	configs.ReadConfig()
}

func main() {
	e := echo.New()
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		Skipper:      echoMiddleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut},
		MaxAge:       86400,
	}))

	db := databases.NewPostgres()

	userRepo := repository.NewUserRepository(db)
	taskRepo := repository.NewTaskRepository(db)

	userService := service.NewUserService(userRepo)
	taskService := service.NewTaskService(taskRepo, userRepo)

	e.GET("/health", func(c echo.Context) error {
		return response.Success(c, map[string]interface{}{
			"time":       time.Now(),
			"ip_address": c.RealIP(),
		})
	})

	userGroup := e.Group("/users", middleware.Middleware())
	handler.NewUserHandler(userGroup, userService)

	taskGroup := userGroup.Group("/:id/tasks", middleware.TaskMiddlerware())
	handler.NewTaskHandler(taskGroup, taskService)

	// Start server
	go func() {
		if err := e.Start(":" + configs.C.Server.Port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatalf("[CRITICAL] Shutting down the server: %+v\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	// SIGINT handles Ctrl+C locally.
	// SIGTERM handles Cloud Run termination signal.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Printf("[CRITICAL] Server shutdown failed: %+v", err)
	}
	e.Logger.Printf("server existed")
}
