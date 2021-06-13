package main

import (
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/manabie-com/togo/internal/handler"
	"github.com/manabie-com/togo/internal/storages/ent"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	client := open("postgresql://ad:pwd@localhost/myapp")

	apiHandler := handler.NewApiHandler("wqGyEBBfPK9w3Lxw", client)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"*"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", handler.ErrorHeader},
		ExposedHeaders:   []string{handler.ErrorHeader},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Post("/login", apiHandler.Login())
	r.Route("/tasks", func(r chi.Router) {
		r.Use(apiHandler.AuthenticationMiddleware)
		r.Get("/", apiHandler.GetTaskByDate())
		r.Post("/", apiHandler.CreateTask())

	})
	log.Info("Server listen on port 5050")
	_ = http.ListenAndServe(":5050", r)
}

func open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}
