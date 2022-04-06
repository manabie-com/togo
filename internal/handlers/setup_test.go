package handlers

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	appPort = 5050
)

var (
	dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		"togo_user",
		"togo_password",
		"localhost",
		"5432",
		"togo_db_test")
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func routes() http.Handler {
	mux := chi.NewRouter()

	{
		mux.Use(middleware.Recoverer)
		mux.Use(DefaultMiddleware)
	}

	{
		mux.Post("/login", Repo.Login)
		mux.Get("/tasks", Repo.RetrieveTasks)
		mux.Post("/tasks", Repo.AddTask)
	}

	return mux
}

func DefaultMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		w.Header().Set("Content-Type", "application/json")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
