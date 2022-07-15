package main

import (
	"log"

	"github.com/datshiro/togo-manabie/internal/infras/app"
)

func main() {
	server := app.NewApp()
	// parse parameter and environment
	server.Parse()

	// config error handler
	server.ConfigErrHandler()

	// config log level
	server.ConfigLogLevel()

	// config log format
	server.ConfigLogFormat()

	//config middleware
	server.ConfigMiddleware()

	// register handlers
	server.RegisterHandlers()

	// run app
	if err := server.Run(); err != nil {
		log.Fatalf("fail to start, err=%v", err)
	}
}
