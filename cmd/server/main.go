package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nvhai245/togo/internal/controller"
	"github.com/nvhai245/togo/internal/repository/task"
	"github.com/nvhai245/togo/internal/service"
)

func main() {
	handler := controller.NewTaskController(
		service.NewTaskService(
			task.NewRepository(
				"sqlite3",
				"./storage/task.repository",
				1,
				1,
			),
		),
	)

	router := gin.Default()
	router.POST("/task", handler.CreateTask)

	if err := router.Run("localhost:8080"); err != nil {
		panic(err)
	}
}
