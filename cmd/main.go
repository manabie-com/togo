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

	storage, err := store.NewStorage(cfg)
	if err != nil {
		log.Fatal("connect to store fail:", logger.Object("error", err))
	}

	svc, err := service.NewTogoService(cfg, storage, log)
	if err != nil {
		log.Fatal("create service fail:", logger.Object("error", err))
	}

	server := NewServer(cfg, svc)
	err = server.start()
	if err != nil {
		log.Fatal("staring server fail:", logger.Object("error", err))
	}
}
