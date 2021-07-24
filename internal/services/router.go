package services

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func NewRouter() *Router {
	r := mux.NewRouter()
	return &Router{
		router: r,
	}
}

type Router struct {
	router *mux.Router
}

func (r *Router) AddHandler(path string, handler func(req *http.Request) (resp interface{}, err error), interceptor httpInterceptor, methods ...string) {
	route := r.router.HandleFunc(path, interceptor.WithHandler(handler))
	if len(methods) > 0 {
		route.Methods(methods...)
	}
}

func (r *Router) Start(port int32) {
	srv := &http.Server{
		Handler:      r.router,
		Addr:         fmt.Sprintf(":%d", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
