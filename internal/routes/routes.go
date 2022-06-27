package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("lntvan166: Hello from Home!"))

}
func HandleRequest(route *mux.Router) {

	route.HandleFunc("/", Home)

	HandleUser(route)
	HandleAuthentication(route)
	HandleTask(route)
	HandlePlan(route)

	http.Handle("/", route)
}
