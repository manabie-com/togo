package routes

import (
	"lntvan166/togo/controller/task"
	"net/http"
)

func TaskRoute(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		task.GetTaskForUser(w, r)
	case "POST":

	case "DELETE":

	default:
		http.ServeFile(w, r, "./views/index.html")
	}
}
