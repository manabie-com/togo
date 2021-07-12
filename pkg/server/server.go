package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/manabie-com/togo/internal/services/auth"
	"github.com/manabie-com/togo/internal/services/task"
	"github.com/manabie-com/togo/internal/storages/database"
	"github.com/manabie-com/togo/pkg/jwtprovider"
	"github.com/manabie-com/togo/pkg/server/config"
	"github.com/manabie-com/togo/pkg/server/handler"
	"github.com/manabie-com/togo/pkg/server/middleware"
	"github.com/manabie-com/togo/server"
)

func Serve() {

	jwtProvider := jwtprovider.NewJWTProvider(config.JWTKey, config.JWTExpiresIn)
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}

	userRepo := database.NewUserRepo(db)
	taskRepo := database.NewTaskRepo(db)

	authService := auth.NewAuthService(userRepo, jwtProvider)
	taskService := task.NewTaskService(taskRepo, userRepo)

	taskHandler := handler.NewTaskHandler(taskService)
	authHandler := handler.NewAuthHandler(authService)

	authMiddleware := middleware.NewAuthMiddleware(jwtProvider)

	router := mux.NewRouter()

	router.HandleFunc("/login", authHandler.Login)
	router.Handle("/tasks", authMiddleware.IsAuthorized(http.HandlerFunc(taskHandler.Create))).Methods(http.MethodPost)
	router.Handle("/tasks", authMiddleware.IsAuthorized(http.HandlerFunc(taskHandler.List))).Methods(http.MethodGet)

	s := server.NewServer(server.NewConfig(5050))
	s.Register(router)
	log.Printf("Starting server... at port: %d", 5050)
	err = s.Serve()
	if err != nil {
		log.Fatalf("Error serve servers cause by %w", err)
	}
}
