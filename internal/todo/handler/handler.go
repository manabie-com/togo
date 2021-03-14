package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/manabie-com/togo/internal/todo/domain"
	s "github.com/manabie-com/togo/internal/todo/service"
)

var tokenAuth *jwtauth.JWTAuth

type TodoRepositoryList struct {
	UserRepo domain.UserRepository
	TaskRepo domain.TaskRepository
}

func init() {
	tokenAuth = jwtauth.New("HS256", []byte(os.Getenv("JWT_KEY")), nil)
}

func NewTodoHandler(todoRepo TodoRepositoryList) http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Protected routes
	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		taskService := s.NewTaskService(todoRepo.TaskRepo)
		taskHandler := NewTaskHandler(taskService)
		r.Get("/tasks", taskHandler.ListTask)
		r.Post("/tasks", taskHandler.CreateTask)
	})

	r.Group(func(r chi.Router) {
		// Public routes
		authService := s.NewAuthService(todoRepo.UserRepo)
		authHandler := NewAuthHandler(authService)
		r.Post("/login", authHandler.Login)
	})

	return r
}

type AppHandler struct{}

func (h *AppHandler) getUserIDFromCtx(r *http.Request) (int, error) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return 0, err
	}

	userID, ok := claims["userID"].(float64)
	if !ok {
		return 0, errors.New("unable to parse user id")
	}

	return int(userID), nil
}

func (h *AppHandler) responseError(w http.ResponseWriter, statusCode int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{
		"error": msg,
	})
}
