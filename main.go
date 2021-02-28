package main

import (
	"database/sql"
	"github.com/manabie-com/togo/config"
	cc "github.com/manabie-com/togo/pkg/common/config"
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/services"

	_ "github.com/lib/pq"
)

func main() {
	cc.InitFlags()
	cc.ParseFlags()

	cfg, err := config.Load()
	if err != nil {
		log.Panicf("error loading config %v", err)
	}

	db, err := sql.Open(cfg.Postgres.ConnectionString())
	if err != nil {
		log.Panicf("error opening db %v", err)
	}

	todoService := services.NewToDoService(db, cfg.JWTSecret, cfg.MaxTodo)

	log.Printf("HTTP server listening at %v", cfg.HTTP.Address())
	err = http.ListenAndServe(cfg.HTTP.Address(), todoService)
	if err != nil {
		log.Panicf("error when starting server %v", err)
	}
}
