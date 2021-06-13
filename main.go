package main

import (
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/gomodule/redigo/redis"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/manabie-com/togo/internal/handler"
	"github.com/manabie-com/togo/internal/storages/ent"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func main() {
	client := getDatabaseClient("postgresql://ad:pwd@localhost/myapp")

	redisPool := getRedisPool("localhost:6379")

	jwtKey := "wqGyEBBfPK9w3Lxw"
	apiHandler := handler.NewApiHandler(jwtKey, client)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"*"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", handler.XErrorMessage, handler.XRateLimit},
		ExposedHeaders:   []string{handler.XErrorMessage, handler.XRateLimit},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Post("/login", apiHandler.Login())
	r.Route("/tasks", func(r chi.Router) {
		r.Use(handler.NewAuthenticator(jwtKey))
		r.Get("/", apiHandler.GetTaskByDate())

		r.With(handler.NewRateLimiter(redisPool)).Post("/", apiHandler.CreateTask())

	})
	log.Info("Server listen on port 5050")
	_ = http.ListenAndServe(":5050", r)
}

func getDatabaseClient(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func getRedisPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}
