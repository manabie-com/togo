package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"task-manage/internal/api"
	db "task-manage/internal/db/sqlc"
	"task-manage/internal/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	queries := db.New(conn)
	server, err := api.NewServer(config, queries)
	err = server.Start(config.HttpServerAddr)
	if err != nil {
		log.Fatal("cannot start application", err)
	}

}
