package controller

import (
	"database/sql"
	"lntvan166/togo/internal/repository"
	"lntvan166/togo/internal/usecase"
	"lntvan166/togo/pkg"
)

type Handler struct {
	UserController
	TaskController
	AuthController
}

var HandlerInstance *Handler

func NewHandler(db *sql.DB) {
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	crypto := pkg.NewCrypto()

	userUsecase := usecase.NewUserUsecase(userRepository, taskRepository, crypto)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, userRepository)

	userController := NewUserController(userUsecase, taskUsecase)
	taskController := NewTaskController(taskUsecase, userUsecase)
	authController := NewAuthController(userUsecase)

	HandlerInstance = &Handler{
		UserController: *userController,
		TaskController: *taskController,
		AuthController: *authController,
	}
}
