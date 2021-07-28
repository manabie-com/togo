package router

import (
	middleware "github.com/manabie-com/togo/internal/middlewares"
	"github.com/manabie-com/togo/internal/transport"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func NewRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.CORS)

	transport := transport.NewTransport(db)
	router.HandleFunc("/login", transport.Login)
	router.HandleFunc("/tasks", middleware.Authorization(transport.CreateTask)).Methods("POST")
	router.HandleFunc("/tasks", middleware.Authorization(transport.ListTasks)).Methods("GET")

	return router
}
