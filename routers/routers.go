package routers

import (
	"github.com/gorilla/mux"
)

func Routing(r *mux.Router){
    userRouter(r)
    taskRouter(r)
    authRouting := r.PathPrefix("/auth").Subrouter()
    authRouting.HandleFunc("/{path}", AuthHandle).Methods("POST")
}