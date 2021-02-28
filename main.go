package main

import (
	"database/sql"
	"github.com/manabie-com/togo/config"
	cc "github.com/manabie-com/togo/pkg/common/config"
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages/sqlstore"

	_ "github.com/lib/pq"
)

func main() {
	cc.InitFlags()
	cc.ParseFlags()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("error loading config", err)
	}

	db, err := sql.Open(cfg.Postgres.ConnectionString())
	if err != nil {
		log.Fatal("error opening db", err)
	}

	http.ListenAndServe(":5050", &services.ToDoService{
		JWTKey: cfg.JWTSecret,
		Store: &sqlstore.Store{
			DB: db,
		},
	})
}
