package app

import (
	"github.com/gorilla/mux"
	"github.com/manabie-com/togo/controllers"
)

func (a *App) Routes() {

	router := a.Router
	// sub router like http://<HOST>:<PORT>/api/users
	userRouter := router.PathPrefix("/api/users").Subrouter()
	userRouter.HandleFunc("/me", a.GetMe).Methods("GET")
	userRouter.HandleFunc("/signup", a.SignUp).Methods("POST")
	userRouter.HandleFunc("/login", a.Login).Methods("POST")
	userRouter.HandleFunc("/edit", a.UpdateMe).Methods("PATCH")
	userRouter.HandleFunc("/delete", a.DeleteMe).Methods("DELETE")
	// sub router like http://<HOST>:<PORT>/api/tasks
	taskRouter := router.PathPrefix("/api/tasks").Subrouter()
	taskRouter.HandleFunc("", a.GetTasks).Methods("GET")
	taskRouter.HandleFunc("/{id}", a.GetTask).Methods("GET")
	taskRouter.HandleFunc("/add", a.Add).Methods("POST")
	taskRouter.HandleFunc("/{id}", a.Edit).Methods("PATCH")
	taskRouter.HandleFunc("/{id}", a.Delete).Methods("DELETE")
	// sub routes like http://<HOST>:<PORT>/api/payments
	paymentRouter := router.PathPrefix("/api/payments").Subrouter()
	paymentRouter.HandleFunc("", a.Payment).Methods("POST")
	// runs database
	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(controllers.JwtAuthentication)
}
