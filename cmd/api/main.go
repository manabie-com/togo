package main

import (
	"github.com/TrinhTrungDung/togo/config"
	"github.com/TrinhTrungDung/togo/pkg/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	// Initialize HTTP server
	e := server.New(&server.Config{
		Port:         cfg.Port,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	})

	// Start the HTTP server
	server.Start(e)
}
