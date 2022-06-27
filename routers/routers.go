package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/huynhhuuloc129/todo/controllers"
)

func Routing(r *mux.Router, h *controllers.BaseHandler){
    userRouter(r, h)
    taskRouter(r, h)
    authRouting := r.PathPrefix("/auth").Subrouter()
    authRouting.HandleFunc("/{path}", func(w http.ResponseWriter, r *http.Request) {
        params := mux.Vars(r)
        path := params["path"]
    
        switch path {
        case "register": // url match register link
            h.Register(w, r)
    
        case "login": // url match login link
            h.Login(w, r)
        default: // not match any link
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    }).Methods("POST")   
}

