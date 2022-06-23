package routers

import (
	"net/http"
	"strconv"

	"github.com/huynhhuuloc129/todo/controllers"
)

func OneTaskHandle(w http.ResponseWriter, r *http.Request, params []string, username string, userid int) { // Handle query for one user
	id, err := strconv.Atoi(params[2])
	if err != nil {
		http.Error(w, "convert string id to int failed", http.StatusBadRequest)
	}

	switch r.Method {
	case http.MethodGet: 						// method GET
		controllers.ResponeOneTask(w, r, id, userid)
	case http.MethodDelete: 					// method DELETE
		controllers.DeleteTask(w, r, id, userid)
	case http.MethodPut: 						// method PUT
		controllers.UpdateTask(w, r, id, userid)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func AllTaskHandle(w http.ResponseWriter, r *http.Request, username string, userid int) { // Handle query for all user
	switch r.Method {
	case http.MethodGet: 						// method GET
		controllers.ResponeAllTask(w, r, userid)
	case http.MethodPost: 						// method POST
		controllers.CreateTask(w, r, userid)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
