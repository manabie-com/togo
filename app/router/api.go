package router

import (
	"github.com/gorilla/mux"
	"github.com/huuthuan-nguyen/manabie/app/handler"
	"github.com/huuthuan-nguyen/manabie/app/middleware"
	"github.com/huuthuan-nguyen/manabie/config"
	"net/http"
)

func SetAPIRoutes(router *mux.Router, config *config.Config, handler *handler.Handler) *mux.Router {
	apiRouter := router.PathPrefix("/api").Subrouter()

	// authentication
	apiRouter.HandleFunc("/register", handler.UserRegister).Methods(http.MethodPost, http.MethodOptions)
	apiRouter.HandleFunc("/auth/login", handler.UserLogin).Methods(http.MethodPost, http.MethodOptions)

	// tasks
	productRouter := apiRouter.PathPrefix("/").Subrouter()
	productRouter.HandleFunc("/tasks", handler.TaskStore).Methods(http.MethodPost, http.MethodOptions)
	productRouter.HandleFunc("/tasks/{id:[0-9]+}", handler.TaskUpdate).Methods(http.MethodPut, http.MethodOptions)
	productRouter.HandleFunc("/tasks/{id:[0-9]+}", handler.TaskDestroy).Methods(http.MethodDelete, http.MethodOptions)
	productRouter.HandleFunc("/tasks", handler.TaskIndex).Methods(http.MethodGet, http.MethodOptions)
	productRouter.HandleFunc("/tasks/{id:[0-9]+}", handler.TaskShow).Methods(http.MethodGet, http.MethodOptions)

	jwtAuthenticationMiddleware := middleware.JWTAuthenticateMiddleware(config, handler.GetDB())
	productRouter.Use(jwtAuthenticationMiddleware)

	return router
}
