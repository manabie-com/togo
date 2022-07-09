package http

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lawtrann/togo"
)

// Server represents an HTTP server. It is meant to wrap all HTTP functionality used by the application
// so that dependent packages (such as cmd/webservice) don't need to reference the "net/http" package at all.
type Server struct {
	// A Server defines parameters for running an HTTP server.
	Server *http.Server
	// Mux is a simple HTTP route multiplexer that parses a request path, records any URL params, and executes an end handler
	Router *chi.Mux

	// Servics used by the various HTTP routes.
	TodoService togo.TodoService
}

func NewServer() *Server {
	// Create a new server that wraps the net/http server & add a chi router.
	s := &Server{
		Server: &http.Server{},
		Router: chi.NewRouter(),
	}

	// RequestID is a middleware that injects a request ID into the context of each request.
	s.Router.Use(middleware.RequestID)
	// Logger is a middleware that logs the start and end of each request.
	s.Router.Use(middleware.Logger)
	// Recoverer is a middleware that recovers from panics, logs the panic, and returns a HTTP 500 status if possible.
	s.Router.Use(middleware.Recoverer)
	// Stop processing after 2.5 seconds.
	s.Router.Use(middleware.Timeout(2500 * time.Millisecond))
	// CleanPath middleware will clean out double slash mistakes from a user's request path.
	s.Router.Use(middleware.CleanPath)
	// Allowed content type.
	s.Router.Use(middleware.AllowContentType("application/json"))

	// Setup error handling routes.
	s.Router.NotFound(s.HandleNotFound)
	// Setup method not allowed.
	s.Router.MethodNotAllowed(http.HandlerFunc(s.HandleMethodNotAllowed))

	// Register api routes.
	s.Router.Mount("/api", s.RegisterTodoRoutes())

	return s
}

// Handles requests to routes that don't exist.
func (s *Server) HandleNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	// Encode response as JSON
	resj := togo.TemplateResponse{
		Status:  http.StatusNotFound,
		Message: togo.ErrPageNotFound.Error(),
	}
	err := encodeResponseAsJSON(resj, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

// Handles requests to routes that method is not allowed.
func (s *Server) HandleMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	// Encode response as JSON
	resj := togo.TemplateResponse{
		Status:  http.StatusMethodNotAllowed,
		Message: togo.ErrHTTPMethodNotAllowed.Error(),
	}
	err := encodeResponseAsJSON(resj, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

func (s *Server) ListenAndServe(port string) error {
	return http.ListenAndServe(port, s.Router)
}

// Utility functions
func encodeResponseAsJSON(data interface{}, w io.Writer) error {
	return json.NewEncoder(w).Encode(data)
}
