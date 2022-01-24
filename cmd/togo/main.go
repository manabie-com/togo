package main

import (
	"github.com/sirupsen/logrus"
	"github.com/trinhdaiphuc/togo/internal"
)

func main() {
	server, cleanup, err := internal.InitializeServer()
	if err != nil {
		panic(err)
	}
	defer func() {
		logrus.Info("Running cleanup tasks...")
		cleanup()
	}()
	server.Run()
}
