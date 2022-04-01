package main

import (
	"fmt"

	"github.com/TrinhTrungDung/togo/config"
	"github.com/TrinhTrungDung/togo/pkg/db"
	"github.com/TrinhTrungDung/togo/pkg/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	db, err := db.New(fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=%s", cfg.DbDialect, cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName, cfg.DbSslMode), cfg.DbLog)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Initialize HTTP server
	e := server.New(&server.Config{
		Port:         cfg.ServerPort,
		ReadTimeout:  cfg.ServerReadTimeout,
		WriteTimeout: cfg.ServerWriteTimeout,
	})

	// Start the HTTP server
	server.Start(e)
}
