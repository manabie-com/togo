package main

import (
	"github.com/gin-gonic/gin"
	db "github.com/qgdomingo/todo-app/database"
	"github.com/qgdomingo/todo-app/controller"
	"github.com/qgdomingo/todo-app/repository"
	"fmt"
	"os"
)

func main() {
	dbPoolConn, err := db.CreateConnection()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer dbPoolConn.Close()

	// initialize the repository
	taskRepo := repository.TaskRepository{ DBPoolConn: dbPoolConn }

	// create controllers for both task and user
	taskCont := controller.TaskController{ TaskRepo: &taskRepo }

	router := gin.Default()

	router.POST("/login", func(context *gin.Context){})
	router.POST("/register", func(context *gin.Context){})

	todoGroup := router.Group("/todo") 
	{
		todoGroup.GET("/fetch", taskCont.GetTasks)
		todoGroup.GET("/fetch/:id", taskCont.GetTaskById)
		todoGroup.GET("/fetch/usertask/:user", taskCont.GetTaskByUser)
		todoGroup.POST("/create", taskCont.CreateTask)
		todoGroup.PUT("/update/:id", taskCont.UpdateTask)
		todoGroup.DELETE("/delete/:id", taskCont.DeleteTask)
	}

	userGroup := router.Group("/user") 
	{
		userGroup.POST("/fetch/:username", func(context *gin.Context){})
		userGroup.POST("/update/:username", func(context *gin.Context){})
	}

	// listen and serve on localhost:8080
	router.Run()
}