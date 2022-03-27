package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/manabie-com/togo/internal/handlers"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	// Use handy middlewares
	{
		mux.Use(middleware.Recoverer)
		mux.Use(DefaultMiddleWare)
	}

	// New api should be added below
	{
		mux.Post("/login", handlers.Repo.Login)
		mux.Get("/tasks", handlers.Repo.RetrieveTasks)
		mux.Post("/tasks", handlers.Repo.AddTask)
	}

	return mux
}
