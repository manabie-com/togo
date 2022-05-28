package main

import (
	"github.com/dinhquockhanh/togo/internal/app/api"
	"github.com/dinhquockhanh/togo/internal/pkg/config"
	"github.com/dinhquockhanh/togo/internal/pkg/http/server"
	"github.com/dinhquockhanh/togo/internal/pkg/log"
)

func main() {
	config.Load()
	log.Init(log.Fields{
		"service": "togo",
	})

	log.Infof("%+v\n", config.All)
	handler, err := api.NewHandler()
	if err != nil {
		log.Panicf("init handler %v", err)

	}
	router, err := api.NewRouter(handler)
	if err != nil {
		log.Panicf("init route %v", err)
	}
	server.Start(&config.All.Server, router)
}
