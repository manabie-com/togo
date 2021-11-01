package api

import (
	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
	"github.com/quochungphp/go-test-assignment/src/domain/auth"
	"github.com/quochungphp/go-test-assignment/src/domain/tasks"
	"github.com/quochungphp/go-test-assignment/src/domain/users"
	"github.com/quochungphp/go-test-assignment/src/infrastructure/middlewares"
)

type APIs struct {
	Router    *mux.Router
	PgSession *pg.DB
}

// Init ...
func (A APIs) Init() {
	router := A.Router
	pgSession := A.PgSession
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
}
