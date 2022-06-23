package routers

import (
	"net/http"
	"strings"

	"github.com/huynhhuuloc129/todo/jwt"
)

const (
	userURL = "users"
	authURL = "auth"
	taskURL = "tasks"
)

func Handle(w http.ResponseWriter, r *http.Request) { // routing for different route
	params := strings.Split(r.RequestURI, "/")
	switch params[1] {
	case userURL: // match user url
		if len(params) > 2 {
			OneUserHandle(w, r, params)
		} else {
			AllUserHandle(w, r)
		}
	case taskURL: 
		username,id,  ok := jwt.CheckToken(w, r)
		if ok {
			if len(params) > 2 {
				OneTaskHandle(w, r, params, username, id)
			} else {
				AllTaskHandle(w, r, username, id)
			}
		} else {
			http.Error(w, "you need to get login token to perform this action", http.StatusNetworkAuthenticationRequired)
		}
	case authURL: //match auth url
		AuthHandle(w, r, params)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
