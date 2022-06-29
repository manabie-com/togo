package routers

import (
	"github.com/gorilla/mux"
	"github.com/huynhhuuloc129/todo/controllers"
	"github.com/huynhhuuloc129/todo/middlewares"
)

func Routing(r *mux.Router, bh *controllers.BaseHandler){
    userRouter(r, bh)
    taskRouter(r, bh)
    authRoutingLogin := r.PathPrefix("/auth/login").Subrouter()
    authRoutingLogin.HandleFunc("", bh.Login).Methods("POST")

    authRoutingRegister := r.PathPrefix("/auth/register").Subrouter()
    authRoutingRegister.HandleFunc("", middlewares.ValidUsernameAndHashPassword(bh ,bh.Register)).Methods("POST")
}

