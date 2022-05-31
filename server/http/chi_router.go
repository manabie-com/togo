package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type chiRouter struct{}

var (
	chiDispatcher = chi.NewRouter()
)

func NewChiRouter() Router {
	return &chiRouter{}
}

func (r *chiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Get(uri, f)
}
func (r *chiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Post(uri, f)
}
func (r *chiRouter) SERVE(port string) {
	port = fmt.Sprintf(":%v", port)
	http.ListenAndServe(port, chiDispatcher)
}
