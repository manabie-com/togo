package routes

import (
	"lntvan166/togo/internal/delivery"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("lntvan166: Hello from Home!"))

}
func HandleRequest(route *mux.Router, handler *delivery.Handler) {

	route.HandleFunc("/", Home)

	HandleUser(route, handler)
	HandleAuthentication(route, handler)
	HandleTask(route, handler)
	HandlePlan(route, handler)

	http.Handle("/", route)
}
