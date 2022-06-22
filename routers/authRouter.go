package routers

import (
	"net/http"

	"github.com/huynhhuuloc129/todo/controllers"
)

func AuthHandle(w http.ResponseWriter, r *http.Request, params []string) { // Handle different request
	switch params[2] {
	case "register":
		switch r.Method {
		case http.MethodPost:
			controllers.Register(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	case "login":
		switch r.Method {
		case http.MethodPost:
			controllers.Login(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}
