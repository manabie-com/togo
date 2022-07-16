package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"manabie/todo/api/settings"
	"manabie/todo/api/tasks"
	"manabie/todo/api/users"
	"manabie/todo/pkg/db"

	settingRespository "manabie/todo/repository/setting"
	taskRespository "manabie/todo/repository/task"
	userRespository "manabie/todo/repository/user"
	settingService "manabie/todo/service/setting"
	taskService "manabie/todo/service/task"
	userService "manabie/todo/service/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Health check
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	{
		// Start Database
		if err := db.Setup(); err != nil {
			panic(err)
		}
	}

	// Repository
	userRp := userRespository.NewUserRespository()
	taskRp := taskRespository.NewTaskRespository()
	settingRp := settingRespository.NewSettingRespository()

	// Service
	userSv := userService.NewUserService(userRp)
	settingSv := settingService.NewSettingService(settingRp)
	taskSv := taskService.NewTaskService(taskRp, settingRp)

	// Handler
	users.NewUserHandler(e, userSv)
	tasks.NewTaskHandler(e, taskSv)
	settings.NewSettingHandler(e, settingSv)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))

	go func() {
		e.Logger.Fatal(e.Start(":8080"))
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	e.Logger.Fatal(e.Shutdown(ctx))
}
