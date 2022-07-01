package routes

import (
	"TOGO/controllers"
	"TOGO/middleware"

	"github.com/gorilla/mux"
)

func UserRoute(router *mux.Router) {
	router.HandleFunc("/user/{Id}", controllers.GetUser()).Methods("GET")
	router.HandleFunc("/user", middleware.AuthMiddleware(controllers.UpdateMe())).Methods("PUT")
	router.HandleFunc("/user/{Id}", middleware.AuthMiddleware(controllers.DeleteUser())).Methods("DELETE")
	//----------------------------------------------------------------
	router.HandleFunc("/user/signup", controllers.Signup()).Methods("POST")
	router.HandleFunc("/users", middleware.AuthMiddleware(controllers.GetAllUser())).Methods("GET")
	router.HandleFunc("/user/login", controllers.Login()).Methods("POST")
	router.HandleFunc("/me", middleware.AuthMiddleware(controllers.GetMe())).Methods("GET")
	router.HandleFunc("/limit", middleware.AuthMiddleware(controllers.UpdateLimit())).Methods("PUT")
}
