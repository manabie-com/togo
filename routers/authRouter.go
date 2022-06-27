package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/huynhhuuloc129/todo/controllers"
)

const (
	registerURL = "register"
	loginURL    = "login"
)

// Handle different request
func AuthHandle(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	path := params["path"]

	switch path {
	case registerURL: // url match register link
		controllers.Register(w, r)

	case loginURL: // url match login link
		controllers.Login(w, r)
	default: // not match any link
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
