package routers

import (
	"net/http"

	"github.com/huynhhuuloc129/todo/controllers"
)

const (
	registerURL = "register"
	loginURL = "login"
)

func AuthHandle(w http.ResponseWriter, r *http.Request, params []string) { // Handle different request
	switch params[2] {
	case registerURL: 			// url match register link
		switch r.Method {
		case http.MethodPost: 	// and match method POST (ONLY METHOD POST)
			controllers.Register(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	case loginURL: 				// url match login link
		switch r.Method {
		case http.MethodPost:	// and match method POST (ONLY METHOD POST)
			controllers.Login(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	default: 					// not match any link
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}
