package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	mm "github.com/manabie-com/togo/internal/pkg/middleware"
	"github.com/manabie-com/togo/internal/todo/domain"
	s "github.com/manabie-com/togo/internal/todo/service"
	log "github.com/sirupsen/logrus"
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

	// Http log
	logger := log.New()
	logger.Formatter = &log.JSONFormatter{
		DisableTimestamp: true,
	}
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(mm.NewStructuredLogger(logger))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5))
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middleware.Heartbeat("/ping"))
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

func (h *AppHandler) getUserIDFromCtx(ctx context.Context) (int, error) {
	_, claims, err := jwtauth.FromContext(ctx)
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
