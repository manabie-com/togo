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

	_, err = db.New(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbPort, cfg.DbSslMode), cfg.DbLog)
	if err != nil {
		panic(err)
	}

	// Initialize HTTP server
	e := server.New(&server.Config{
		Port:         cfg.ServerPort,
		ReadTimeout:  cfg.ServerReadTimeout,
		WriteTimeout: cfg.ServerWriteTimeout,
	})

	// Start the HTTP server
	server.Start(e)
}
