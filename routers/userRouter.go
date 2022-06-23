package routers

import (
	"net/http"
	"strconv"

	"github.com/huynhhuuloc129/todo/controllers"
)

func OneUserHandle(w http.ResponseWriter, r *http.Request, params []string) { // Handle query for one user
	id, err := strconv.Atoi(params[2])
	if err != nil {
		http.Error(w, "convert string id to int failed", http.StatusBadRequest)
	}

	switch r.Method { 		// handle for one user
	case http.MethodGet: 	// method GET
		controllers.ResponeOneUser(w, r, id)
	case http.MethodDelete: // method DELETE
		controllers.DeleteUser(w, r, id)
	case http.MethodPut: 	// method PUT
		controllers.UpdateUser(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func AllUserHandle(w http.ResponseWriter, r *http.Request) { // Handle query for all user
	switch r.Method {
	case http.MethodGet: 	// method GET
		controllers.ResponeAllUser(w, r)
	case http.MethodPost:	// method POST
		controllers.CreateUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}