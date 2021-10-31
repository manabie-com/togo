package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/tylerb/graceful"

	"github.com/quochungphp/go-test-assignment/src/domain/auth"
	"github.com/quochungphp/go-test-assignment/src/domain/tasks"
	"github.com/quochungphp/go-test-assignment/src/domain/users"
	"github.com/quochungphp/go-test-assignment/src/infrastructure/middlewares"
	"github.com/quochungphp/go-test-assignment/src/infrastructure/pg_driver"
	"github.com/quochungphp/go-test-assignment/src/infrastructure/redis_driver"
	"github.com/quochungphp/go-test-assignment/src/pkgs/settings"
)

func main() {
	// Init redis
	redisHost := os.Getenv(settings.RedisHost) + ":" + os.Getenv(settings.RedisPort)
	redisDriver := redis_driver.RedisDriver{}
	err := redisDriver.Setup(redis_driver.RedisConfiguration{
		Addr: redisHost,
	})
	if err != nil {
		panic(fmt.Sprint("Error while setup redis driver: ", err))
	}

	// Init postgresql
	pgSession, err := pg_driver.Setup(pg_driver.DBConfiguration{
		Driver:   os.Getenv(settings.DbDriver),
		Host:     os.Getenv(settings.PgHost),
		Port:     os.Getenv(settings.PgPort),
		Database: os.Getenv(settings.PgDB),
		User:     os.Getenv(settings.PgUser),
		Password: os.Getenv(settings.PgPass),
	})
	if err != nil {
		panic(fmt.Sprint("Error while setup postgres driver: ", err))
	}
	// Init Gorilla Router
	router := mux.NewRouter()

	// User create action
	userCreatAction := users.UserCreateAction{pgSession}
	userCtrl := users.UserController{
		userCreatAction,
	}
	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("", userCtrl.Create).Methods("POST").Name("UserCreateAction")

	// Login & logout action
	authLoginAction := auth.AuthLoginAction{pgSession}
	authLogoutAction := auth.AuthLogoutAction{}
	authCtrl := auth.AuthController{
		authLoginAction,
		authLogoutAction,
	}

	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("", authCtrl.Login).Methods("POST").Name("AuthLoginAction")

	authLogoutRouter := router.PathPrefix("/auth").Subrouter()
	authLogoutRouter.HandleFunc("", authCtrl.Logout).Methods("DELETE").Name("AuthLogoutAction")
	authLogoutRouter.Use(middlewares.AuthMiddleware)

	// Create task action
	createTaskAction := tasks.TaskCreateAction{pgSession}
	listTaskAction := tasks.TaskListAction{pgSession}
	taskCtrl := tasks.TaskController{
		createTaskAction,
		listTaskAction,
	}

	taskRouter := router.PathPrefix("/tasks").Subrouter()
	taskRouter.HandleFunc("", taskCtrl.Create).Methods("POST").Name("TaskCreateAction")
	taskRouter.HandleFunc("", taskCtrl.List).Methods("GET").Name("ListTaskAction")
	taskRouter.Use(middlewares.AuthMiddleware)

	srv := &graceful.Server{
		Timeout: 5 * time.Second,
		BeforeShutdown: func() bool {
			pgSession.Close()
			log.Println("Server is shutting down database connection")

			return true
		},
		Server: &http.Server{
			Addr:    ":" + os.Getenv(settings.Port),
			Handler: router,
		},
	}

	log.Printf("Server is runing port: %v\n", os.Getenv(settings.Port))
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("While start server")
	}
}
