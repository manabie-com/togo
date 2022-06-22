package routers

import (
	"net/http"
	"strconv"

	"github.com/huynhhuuloc129/todo/controllers"
)

func OneUserHandle(w http.ResponseWriter, r *http.Request, params []string) { // Handle query for one user
	id, err := strconv.Atoi(params[2])
	controllers.ErrorHandle(w, err, http.StatusMethodNotAllowed)

	switch r.Method {
	case http.MethodGet:
		controllers.ResponeOneUser(w, r, id)
	case http.MethodDelete:
		controllers.DeleteUser(w, r, id)
	case http.MethodPut:
		controllers.UpdateUser(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func AllUserHandle(w http.ResponseWriter, r *http.Request) { // Handle query for all user
	switch r.Method {
	case http.MethodGet:
		controllers.ResponeAllUser(w, r)
	case http.MethodPost:
		controllers.CreateUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}