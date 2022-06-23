package routes

import (
	"lntvan166/togo/controller/auth"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func HandleAuthentication(route *mux.Router) {
	authRouter := route.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/{method}", AuthRoute)
}

func AuthRoute(w http.ResponseWriter, r *http.Request) {
	args := strings.Split(r.URL.Path, "/")

	switch args[2] {
	case "register":
		auth.Register(w, r)
	case "login":
		auth.Login(w, r)
	default:
		w.Write([]byte("lntvan166: invalid authentication method"))
	}

}
