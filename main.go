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

	// create db connection that returns the address of the controller instances
	// use those controller instances as the function handler below
	// db struct should be found in the controllers?
	
	todoCont := taskController.TaskDB{ DBPoolConn: dbPoolConn }

	router := gin.Default()

	todoGroup := router.Group("/todo") 
	{
		todoGroup.GET("/fetch", todoCont.GetTasks)
		todoGroup.GET("/fetch/:id", todoCont.GetTaskById)
		todoGroup.GET("/fetch/usertask/:user", todoCont.GetTaskByUser)
		todoGroup.POST("/create", todoCont.CreateTask)
		todoGroup.PUT("/update/:id", todoCont.UpdateTask)
		todoGroup.DELETE("/delete/:id", todoCont.DeleteTask)
	}

	//userGroup := router.Group("/user") 
	//{
	//	todoGroup.POST("/login", func(context *gin.Context){})
	//}

	// listen and serve on localhost:8080
	router.Run() 
}