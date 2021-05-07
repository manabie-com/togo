package main

import (
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/services"
	"log"
)

func main() {
	cfg := config.Load()
	api, err := services.NewAPI(cfg)
	if err != nil {
		log.Fatal(err)
	}
	api.Start()
}
