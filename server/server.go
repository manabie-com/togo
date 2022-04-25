package server

import (
	"log"
	"net/http"

	"github.com/SVincentTran/togo/handlers"
	"github.com/SVincentTran/togo/middlewares"
	"github.com/gorilla/mux"
)

type Server struct {
	handlers *handlers.Handler
}

func (s Server) Start() {
	r := mux.NewRouter()
	r.Handle("/user/{userId}/todo", middlewares.ErrHandler(s.handlers.TodoTasksHanlder)).Methods(http.MethodPost)
	log.Print("Starting server, listening at port 9000...")
	err := http.ListenAndServe("localhost:9000", r)
	if err != nil {
		log.Panicf("Can not serve server...")
	}
}

func New() *Server {
	handlers := handlers.New()
	return &Server{
		handlers: handlers,
	}
}
