package routers

import (
	"github.com/gorilla/mux"
	"github.com/huynhhuuloc129/todo/controllers"
	"github.com/huynhhuuloc129/todo/middlewares"
)

func Routing(r *mux.Router){
    userRouter(r)
    taskRouter(r)
    authRouting := r.PathPrefix("/auth").Subrouter()
    authRouting.HandleFunc("/{path}", AuthHandle).Methods("POST")
}

func userRouter(r *mux.Router){
    userRouting := r.PathPrefix("/users").Subrouter()
    userRoutingid := r.PathPrefix("/users/{id}").Subrouter()

    userRouting.HandleFunc("/", controllers.ResponeAllUser).Methods("GET")
	userRouting.HandleFunc("/", controllers.CreateUser).Methods("POST")
    
    userRoutingid.Use(middlewares.MiddlewareID)
    userRoutingid.HandleFunc("/", controllers.ResponeOneUser).Methods("GET")
	userRoutingid.HandleFunc("/", controllers.DeleteUser).Methods("DELETE")
    userRoutingid.HandleFunc("/", controllers.UpdateUser).Methods("PUT")
    
   
}
func taskRouter(r *mux.Router){
    taskRouting := r.PathPrefix("/tasks").Subrouter()
    taskRoutingid := r.PathPrefix("/tasks/{id}").Subrouter()

    taskRouting.Use(middlewares.Logging)
    taskRouting.HandleFunc("/", controllers.ResponeAllTask).Methods("GET")
    taskRouting.HandleFunc("/", controllers.CreateTask).Methods("POST")

    taskRoutingid.Use(middlewares.Logging, middlewares.MiddlewareID)
	taskRoutingid.HandleFunc("/", controllers.ResponeOneTask).Methods("GET")
    taskRoutingid.HandleFunc("/", controllers.DeleteTask).Methods("DELETE")
    taskRoutingid.HandleFunc("/", controllers.UpdateTask).Methods("PUT")
}