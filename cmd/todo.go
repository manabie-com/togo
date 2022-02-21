package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	db2 "github.com/chi07/todo/db"
	"github.com/chi07/todo/internal/http/handler"
	"github.com/chi07/todo/internal/repo"
	"github.com/chi07/todo/internal/service"
)

func main() {
	viper.SetEnvPrefix("TODO")
	viper.AutomaticEnv()
	dbURL := viper.GetString("DB_URL")

	logrus.Info("DB_URL: ", dbURL)

	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		logrus.Fatal("Could not connect to database")
	}

	if err := db2.Migrate("file://../db/migrations", dbURL); err != nil {
		logrus.Fatal("Migrate failed: ", err.Error())
	}

	defer func() {
		if err = db.Close(); err != nil {
			log.Error().Err(err).Msg("Cannot close DB connection")
		}
	}()

	taskRepo := repo.NewTask(db)
	limitedRepo := repo.NewLimitation(db)
	CreateTaskService := service.NewCreateTaskService(taskRepo, limitedRepo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/todos", func(r chi.Router) {
		r.Method("POST", "/", handler.NewCreateTaskHandler(CreateTaskService))
	})

	err = http.ListenAndServe(":8081", r)
	log.Error().Err(err).Msg("Stopped serving HTTP")
}
