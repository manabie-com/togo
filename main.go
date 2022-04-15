package main

import (
	"github.com/gin-gonic/gin"
	db "github.com/qgdomingo/todo-app/database"
	taskController "github.com/qgdomingo/todo-app/controller"
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

	// create for user and the login and register controllers
	taskCont := taskController.TaskDB{ DBPoolConn: dbPoolConn }

	router := gin.Default()

	//	todoGroup.POST("/login", func(context *gin.Context){})
	//	todoGroup.POST("/register", func(context *gin.Context){})
	todoGroup := router.Group("/todo") 
	{
		todoGroup.GET("/fetch", taskCont.GetTasks)
		todoGroup.GET("/fetch/:id", taskCont.GetTaskById)
		todoGroup.GET("/fetch/usertask/:user", taskCont.GetTaskByUser)
		todoGroup.POST("/create", taskCont.CreateTask)
		todoGroup.PUT("/update/:id", taskCont.UpdateTask)
		todoGroup.DELETE("/delete/:id", taskCont.DeleteTask)
	}

	//userGroup := router.Group("/user") 
	//{
	//	todoGroup.POST("/fetch/:username", func(context *gin.Context){})
	//	todoGroup.POST("/update/:username", func(context *gin.Context){})
	//}

	// listen and serve on localhost:8080
	router.Run() 
}