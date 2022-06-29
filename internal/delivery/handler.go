package delivery

import (
	"database/sql"
	"lntvan166/togo/internal/repository"
	"lntvan166/togo/internal/usecase"
	"lntvan166/togo/pkg"
)

type Handler struct {
	UserDelivery
	TaskDelivery
	AuthDelivery
}

var HandlerInstance *Handler

func NewHandler(db *sql.DB) *Handler {
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	crypto := pkg.NewCrypto()

	userUsecase := usecase.NewUserUsecase(userRepository, taskRepository, crypto)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, userRepository)

	userDelivery := NewUserDelivery(userUsecase, taskUsecase)
	taskDelivery := NewTaskDelivery(taskUsecase, userUsecase)
	authDelivery := NewAuthDelivery(userUsecase)

	return &Handler{
		UserDelivery: *userDelivery,
		TaskDelivery: *taskDelivery,
		AuthDelivery: *authDelivery,
	}
}

func Setup(db *sql.DB) {
	HandlerInstance = NewHandler(db)
}
