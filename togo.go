package main

import (
	"github.com/trangmaiq/togo/internal/config"
	"github.com/trangmaiq/togo/internal/server"
)

func main() {
	cfg := config.Load()
	server.Start(cfg)
}
