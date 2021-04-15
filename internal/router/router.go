package router

import (
	"database/sql"
	"github.com/gorilla/mux"
	taskStorage "github.com/manabie-com/togo/internal/app/task/storage/rdbms"
	taskDelivery "github.com/manabie-com/togo/internal/app/task/transport/rest"
	taskService "github.com/manabie-com/togo/internal/app/task/usecase"
	userStorage "github.com/manabie-com/togo/internal/app/user/storage/rdbms"
	userDelivery "github.com/manabie-com/togo/internal/app/user/transport/rest"
	userService "github.com/manabie-com/togo/internal/app/user/usecase"
	"net/http"
)

func NewRouter(db *sql.DB, dbDriverName string, jwtKey string) *mux.Router {
	userSrv := userService.NewAuthService(userStorage.New(db, dbDriverName), jwtKey)
	userHandler := userDelivery.NewDelivery(userSrv)
	taskSrv := taskService.NewTaskService(taskStorage.NewTaskStorage(db, dbDriverName))
	taskHandler := taskDelivery.NewDelivery(taskSrv)

	r := mux.NewRouter()
	r.Path("/login").Methods(http.MethodGet).HandlerFunc(userHandler.Login)
	r.Path("/tasks").Methods(http.MethodPost).HandlerFunc(taskHandler.AddTask)
	r.Path("/tasks").Methods(http.MethodGet).HandlerFunc(taskHandler.RetrieveTasks)

	// middlewares
	r.Use(userHandler.Authorize)
	return r
}
