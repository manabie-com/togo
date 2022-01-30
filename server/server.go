package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/manabie-com/togo/pkg/httpx"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/manabie-com/togo/registry"
	"github.com/rs/cors"
)

// Server ...
type Server struct {
	taskDomain *TaskDomain
	userDomain *UserDomain
}

func New(r *registry.Registry) *Server {
	return &Server{
		taskDomain: newTaskDomain(r),
		userDomain: newUserDomain(r),
	}
}

func (s *Server) router() chi.Router {
	r := chi.NewRouter()
	r.Use(APILoggingMiddleware)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:*",
			"https://localhost:*",
		},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}).Handler)

	r.Route("/api", func(r chi.Router) {
		r.Mount("/task", s.taskDomain.router())
		r.Mount("/user", s.userDomain.router())
	})

	return r
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router().ServeHTTP(w, r)
}

func APILoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := httpx.NewLoggingResponseWriter(w)
		defer func() {
			fmt.Printf(fmt.Sprintf("API infomation: %v [%v] [%v]", r.RequestURI, r.Method, lrw.StatusCode))
		}()
		next.ServeHTTP(lrw, r)
	})
}
