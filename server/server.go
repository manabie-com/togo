package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/manabie-com/togo/config"
	"github.com/manabie-com/togo/internal/services/auth"
	"github.com/manabie-com/togo/internal/services/task"
	"github.com/manabie-com/togo/internal/storages/postgresql"
	"github.com/manabie-com/togo/pkg/jwtprovider"
	"github.com/manabie-com/togo/pkg/server"
	"github.com/manabie-com/togo/server/handler"
	"github.com/manabie-com/togo/server/middleware"
)

func Serve() {

	jwtProvider := jwtprovider.NewJWTProvider(config.JWT.Key, config.JWT.ExpiresIn)
	db := getPostgresConnection(config.PostgreSQL)

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

	s := server.NewServer(server.NewConfig(config.HTTPPort))
	s.Register(router)
	log.Printf("Starting server... at port: %d", config.HTTPPort)
	err := s.Serve()
	if err != nil {
		log.Fatalf("Error serve servers cause by %s", err.Error())
	}
}

func getPostgresConnection(postgresCfg config.PostgreSQLConfig) *sql.DB {
	db, err := sql.Open("postgres", postgresCfg.String())
	if err != nil {
		log.Fatalf("Error when connect postgres cause by %s", err.Error())
	}

	db.SetMaxIdleConns(postgresCfg.MaxIdleConns)
	db.SetMaxOpenConns(postgresCfg.MaxOpenConns)
	db.SetConnMaxLifetime(postgresCfg.ConnMaxLifetime)
	return db
}
