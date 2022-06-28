package main

import (
	"database/sql"
	"fmt"
	"lntvan166/togo/internal/config"
	"lntvan166/togo/internal/controller"
	"lntvan166/togo/internal/repository"
	"lntvan166/togo/internal/routes"
	"lntvan166/togo/internal/usecase"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	config.Load()

	var db *sql.DB
	psqlInfo := config.DATABASE_URL
	if psqlInfo == "" {
		psqlInfo = repository.GetPostgresConnectionString()
	}

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository, taskRepository)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, userRepository)

	userController := controller.NewUserController(userUsecase, taskUsecase)
	taskController := controller.NewTaskController(taskUsecase, userUsecase)
	authController := controller.NewAuthController(userUsecase)

	controller.HandlerInstance = &controller.Handler{
		UserController: *userController,
		TaskController: *taskController,
		AuthController: *authController,
	}

	route := mux.NewRouter()
	routes.HandleRequest(route)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server started!")

}
