package handler

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"togo/domain/controllers"
	"togo/domain/middlewares"
	"togo/domain/services"
)

func NewHTTPServer(
	middleware middlewares.Middleware,
	todoService services.TodoService,
) http.Handler {

	router := chi.NewRouter()
	router.Use(
		middleware.WithCors(),
	)

	router.HandleFunc(
		"/health",
		func(writer http.ResponseWriter, request *http.Request) {
			_, _ = fmt.Fprintf(writer, "OK\n")
		})

	router.Route("/todo", NewTodoHandler(controllers.NewTodoController(todoService)))

	return router
}
