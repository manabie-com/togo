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
	// Create the database connection which will be used by the controllers
	dbPoolConn, err := db.CreateConnection()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer dbPoolConn.Close()

	// Create the repositories 
	taskRepo := repository.TaskRepository{ DBPoolConn: dbPoolConn }
	userRepo := repository.UserRepository{ DBPoolConn: dbPoolConn }

	// Create controllers for both task and user that implements the actual repositories
	// -- Repositories are interfaced so it can be implemented differently 
	//    i.e. in a way not to connect to the database for unit testing
	// -- These controllers handle the following: 
	//		-> Validating the HTTP request url parameter (if any)
	//		-> Validating the HTTP request body to ensure they are populated
	//		-> Provides the appropriate HTTP response based from the results of the repository functions
	taskCont := controller.TaskController{ TaskRepo: &taskRepo }
	userCont := controller.UserController{ UserRepo: &userRepo }

	router := gin.Default()

	router.POST("/login", userCont.LoginUser)
	router.POST("/register", userCont.RegisterUser)

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
		userGroup.GET("/fetch/:username", userCont.FetchUserDetails)
		userGroup.PUT("/update/:username", userCont.UpdateUserDetails)
		userGroup.PUT("/pwdchange/:username", userCont.UserPasswordChange)
	}

	// listen and serve on localhost:8080
	router.Run()
}