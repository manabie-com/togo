package routes

import (
	"TOGO/controllers"
	"TOGO/middleware"

	"github.com/gorilla/mux"
)

func TaskRoute(router *mux.Router) {
	router.HandleFunc("/task/{id}", middleware.AuthMiddleware(controllers.GetOneTask())).Methods("GET")
	router.HandleFunc("/task", middleware.AuthMiddleware(controllers.CreateTask())).Methods("POST")
	router.HandleFunc("/user-tasks", middleware.AuthMiddleware(controllers.GetTask())).Methods("GET")
	router.HandleFunc("/task/{id}", middleware.AuthMiddleware(controllers.DeleteTask())).Methods("DELETE")
	router.HandleFunc("/task/{id}", middleware.AuthMiddleware(controllers.UpdateTask())).Methods("PUT")
	router.HandleFunc("/task/status/{id}", middleware.AuthMiddleware(controllers.UpdateTaskStatus())).Methods("PUT")
	router.HandleFunc("/task-status", middleware.AuthMiddleware(controllers.GetTaskDoing())).Methods("GET")
}
