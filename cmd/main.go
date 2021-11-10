package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/config"
	"github.com/manabie-com/togo/internal/domain"
	"github.com/manabie-com/togo/internal/storages/postgresql"
	"github.com/manabie-com/togo/internal/transport/http"
	"github.com/manabie-com/togo/internal/transport/http/middleware"
	"github.com/manabie-com/togo/pkg/token"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	conn, err := connectDB(cfg.Postgresql)
	if err != nil {
		log.Fatal("cannot connect database", err)
	}
	e := echo.New()
	taskRepository := postgresql.NewTaskPostgresqlRepository(conn)
	userRepository := postgresql.NewUserPostgresqlRepository(conn)
	jtwMaker := token.NewJWTMaker(cfg.ApplicationConfig.TokenSecretKey)
	tokenConfig := domain.NewTokenConfig(cfg.ApplicationConfig.TokenDuration)

	taskUseCase := domain.NewTaskUseCase(taskRepository, userRepository)
	authUseCase := domain.NewAuthUseCase(userRepository, jtwMaker, tokenConfig)
	authMiddleware := middleware.NewAuthMiddleware(jtwMaker)
	http.NewTaskHandler(e, taskUseCase, authMiddleware)
	http.NewAuthHandler(e, authUseCase)

	err = e.Start("0.0.0.0:8000")
	if err != nil {
		log.Fatal("cannot start server")
	}
}

// connectDB to Postgresql
func connectDB(cfg config.DBConfig) (*sql.DB, error) {
	return sql.Open(cfg.Driver, cfg.Source)
}
