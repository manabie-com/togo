package main

import (
	"main/config"
	"main/internal/logger"
	"main/internal/service"
	"main/internal/store"
)

func main() {
	cfg := config.Load()
	log := logger.New()

	server := NewServer(cfg)
	storage, err := store.NewStorage(cfg)
	if err != nil {
		log.Fatal("connect to store fail:", logger.Object("error", err))
	}

	svc, err := service.NewService(cfg, &storage)
	if err != nil {
		log.Fatal("create service fail:", logger.Object("error", err))
	}

	err = server.start(svc)
	if err != nil {
		log.Fatal("staring server fail:", logger.Object("error", err))
	}
}
