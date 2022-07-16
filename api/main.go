package main

import (
	"context"
	"net/http"

	"manabie/todo/api/tasks"
	"manabie/todo/api/users"
	"manabie/todo/pkg/db"

	settingRespository "manabie/todo/repository/setting"
	taskRespository "manabie/todo/repository/task"
	userRespository "manabie/todo/repository/user"
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

		got := db.Config{
			User:     "postgres",
			Password: "password",
			Host:     "postgresql.manabie.todo",
			Database: "todo",
		}

		if err := db.Setup(context.Background(), got); err != nil {
			panic(err)
		}
		e.Logger.Infof("open database")
	}

	// Repository
	userRp := userRespository.NewUserRespository()
	taskRp := taskRespository.NewTaskRespository()
	settingRp := settingRespository.NewSettingRespository()

	// Service
	userSv := userService.NewUserService(userRp)
	taskSv := taskService.NewTaskService(taskRp, settingRp)

	// Handler
	users.NewUserHandler(e, userSv)
	tasks.NewTaskHandler(e, taskSv)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
