package main

import (
	"github.com/sirupsen/logrus"
	"github.com/trangmaiq/togo/internal/config"
	"github.com/trangmaiq/togo/internal/registry"
	"github.com/trangmaiq/togo/internal/server"
)

func main() {
	cfg := config.Load()

	err := registry.Init(cfg)
	if err != nil {
		logrus.WithField("err", err).Fatal("init service failed")
	}

	server.StartWithGracefulShutdown(cfg)
}
