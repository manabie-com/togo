package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	handler2 "github.com/manabie-com/togo/internal/handler"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	handler := handler2.NewApiHandler("wqGyEBBfPK9w3Lxw")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/login", handler.Login())
	r.Route("/tasks", func(r chi.Router) {
		r.Use(handler.AuthenticationMiddleware)
		r.Get("/", handler.GetTaskByDate())
		r.Post("/", handler.CreateTask())

	})
	log.Info("Server listen on port 5050")
	_ = http.ListenAndServe(":5050", r)
}
