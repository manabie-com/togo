package main

import (
	"github.com/google/martian/log"
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/services"
)

func main() {
	cfg := config.Load()
	log.Infof("ToGo is running as environment %s", cfg.Environment)
	api, err := services.NewAPI(cfg)
	if err != nil {
		panic(err)
	}
	api.Start()
}
