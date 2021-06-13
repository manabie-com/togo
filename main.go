package main

import (
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httptracer"
	"github.com/gomodule/redigo/redis"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/manabie-com/togo/internal/handler"
	"github.com/manabie-com/togo/internal/storages/ent"
	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"net/http"
	"time"
)

// @title Swagger Togo Service
// @version 0.0.2
// @description This is a Togo service server
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @Ba

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost
// @BasePath /
func main() {
	client := getDatabaseClient("postgresql://ad:pwd@localhost/myapp")

	redisPool := getRedisPool("localhost:6379")

	jwtKey := "wqGyEBBfPK9w3Lxw"
	apiHandler := handler.NewApiHandler(jwtKey, client)

	tracer, _ := initTracer("togo-service")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(httptracer.Tracer(tracer, httptracer.Config{
		ServiceName:    "togo-service",
		ServiceVersion: "v0.0.2",
	}))
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
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:5050/swagger/doc.json"), //The url pointing to API definition"
	))

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

func initTracer(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		log.Error("cannot init Jaeger: %v", err)
	}
	return tracer, closer
}
