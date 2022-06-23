package routes

import (
	"lntvan166/togo/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("lntvan166: Hello from Home!"))

}
func HandleRequest(route *mux.Router) {

	route.HandleFunc("/", Home)

	route.HandleFunc("/user", UserRoute)

	route.Handle("/task", middleware.Authorization(http.HandlerFunc(TaskRoute)))

	route.HandleFunc("/auth/{method}", AuthRoute)

	http.Handle("/", route)
}
