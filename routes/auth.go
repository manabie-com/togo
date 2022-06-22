package routes

import (
	"lntvan166/togo/controller/auth"
	"net/http"
	"strings"
)

func AuthRoute(w http.ResponseWriter, r *http.Request) {
	args := strings.Split(r.URL.Path, "/")
	switch r.Method {
	case "GET":
		w.Write([]byte("lntvan166: auth"))
	case "POST":
		switch args[2] {
		case "register":
			auth.Register(w, r)
		case "login":
			auth.Login(w, r)
		default:
			w.Write([]byte("lntvan166: auth"))
		}
	default:
		http.ServeFile(w, r, "./views/index.html")
	}
}
