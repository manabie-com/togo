package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"togo/configs"
	"togo/handler"
	middle "togo/middleware"
	"togo/pkg/databases"
	"togo/pkg/logger"
	"togo/repository"
	"togo/response"
	"togo/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	configs.ReadConfig()
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut},
		MaxAge:       86400,
	}))

	db := databases.NewPostgres()
	log := logger.NewLogger()
	logger.L = log

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

	taskGroup := e.Group("/tasks", middle.Middleware())
	handler.NewTaskHandler(taskGroup, taskService)

	userGroup := e.Group("/users", middle.Middleware())
	handler.NewUserHandler(userGroup, userService)

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
		e.Logger.Printf("[CRITICAL]  Server shutdown failed: %+v", err)
	}
	e.Logger.Printf("server existed")
}
