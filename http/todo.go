package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/lawtrann/togo"
)

// RegisterTodoRoutes is a helper function for registering all todo routes.
func (s *Server) RegisterTodoRoutes() *chi.Mux {
	r := chi.NewRouter()

	// API endpoint for creating user if not existed
	r.Route("/{username}", func(r chi.Router) {
		r.Use(s.UserCtx)
		// API endpoint for creating todo.
		r.Post("/todos", s.HandlerTodoAdd)
	})

	return r
}

// This struct represent a key/value for username path parameter
type userName struct{}

func (s *Server) UserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")
		// attach username to context
		ctx := context.WithValue(r.Context(), userName{}, username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *Server) HandlerTodoAdd(w http.ResponseWriter, r *http.Request) {
	// Parse
	var todo *togo.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil || todo.Description == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		// Encode response as JSON
		resj := togo.TemplateResponse{
			Status:  http.StatusUnprocessableEntity,
			Message: togo.ErrCouldNotParseObject.Error(),
		}
		err := encodeResponseAsJSON(resj, w)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		return
	}

	userName, ok := r.Context().Value(userName{}).(string)
	if !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Could not parse username"))
		return
	}

	// Handle with TodoService
	res, err := s.TodoService.Add(r.Context(), todo, userName)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// Encode response as JSON
		resj := togo.TemplateResponse{
			Status:  http.StatusOK,
			Message: err.Error(),
		}
		err = encodeResponseAsJSON(resj, w)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		return
	}

	// Render output to the client based on JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	resj := togo.TemplateResponse{
		Status:  http.StatusCreated,
		Message: togo.ErrSuccessAddingNewTodo.Error(),
		Data:    *res,
	}
	err = encodeResponseAsJSON(resj, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}
