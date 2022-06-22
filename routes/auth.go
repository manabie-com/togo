package routes

import (
	"lntvan166/togo/controller/auth"
	"net/http"
)

func AuthRoute(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("lntvan166: auth"))
	case "POST":
		if r.URL.Path == "/auth/register" {
			auth.Register(w, r)
		}
	default:
		http.ServeFile(w, r, "./views/index.html")
	}
}
