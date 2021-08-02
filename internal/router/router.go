package router

import (
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/controllers"
	"github.com/manabie-com/togo/internal/middleware"

	"github.com/gorilla/mux"
)

var r = mux.NewRouter()

func init() {
	log.Println("Initializing Router")
	apiV1 := r.PathPrefix("/").Subrouter()

	//For root functions
	apiV1.HandleFunc("/", RootRoute).Methods("GET")
	apiV1.HandleFunc("/login", controllers.LoginHandler).Methods("POST")
	apiV1.HandleFunc("/register", controllers.CreateAccountHandler).Methods("POST")
	apiV1.HandleFunc("/profile", middleware.UserAuthentication(controllers.ViewProfileHandler)).Methods("GET")

	apiV1.HandleFunc("/tasks", middleware.UserAuthentication(controllers.CreateTaskHandler)).Methods("POST")
	apiV1.HandleFunc("/tasks", middleware.UserAuthentication(controllers.GetTasksHandler)).Methods("GET")

}

func GetRouter() *mux.Router {
	return r
}

func RootRoute(w http.ResponseWriter, r *http.Request) {
	config.ResponseWithSuccess(w, "OK", "Welcome to this API.")
}
