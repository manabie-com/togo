package routers

import (
	"net/http"
	"strings"
)

const (
	userURL = "users"
	authURL = "auth"
)

func Handle(w http.ResponseWriter, r *http.Request) { // routing for different route
	params := strings.Split(r.RequestURI, "/")
	switch params[1] {
	case userURL:
		if len(params) > 2 {
			OneUserHandle(w, r, params)
		} else {
			AllUserHandle(w, r)
		}
	case authURL:
		AuthHandle(w, r, params)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
