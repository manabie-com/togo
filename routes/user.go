package routes

import (
	"lntvan166/togo/controller/user"
	"net/http"
)

func UserRoute(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		user.GetAllUsers(w, r)
	case "POST":
		http.ServeFile(w, r, "./views/index.html")
	default:
		http.ServeFile(w, r, "./views/index.html")
	}
}
