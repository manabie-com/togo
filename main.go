package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/manabie-com/togo/controllers"
)

func main() {
	router := mux.NewRouter()

	// sub router like http://<HOST>:<PORT>/api/users
	userRouter := router.PathPrefix("/api/users").Subrouter()
	userRouter.HandleFunc("/", controllers.GetMe).Methods("GET")
	userRouter.HandleFunc("/signUp", controllers.SignUp).Methods("POST")
	userRouter.HandleFunc("/login", controllers.Login).Methods("POST")
	userRouter.HandleFunc("/", controllers.UpdateMe).Methods("PATCH")
	// sub router like http://<HOST>:<PORT>/api/tasks
	taskRouter := router.PathPrefix("/api/tasks").Subrouter()
	taskRouter.HandleFunc("/", controllers.GetTasks).Methods("GET")
	taskRouter.HandleFunc("/{id}", controllers.GetTask).Methods("GET")
	taskRouter.HandleFunc("/add", controllers.Add).Methods("POST")
	taskRouter.HandleFunc("/{id}", controllers.Edit).Methods("PATCH")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Println(port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
