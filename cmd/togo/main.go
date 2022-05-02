package main

import (
	"github.com/sirupsen/logrus"
	"log"
	"todo/internal"
)

func main() {
	server, cleanup, err := internal.InitializeServer()
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		logrus.Info("Running cleanup tasks...")
		cleanup()
	}()
	server.Run()
}
