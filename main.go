package main

import (
	"os"
	"todo-api/app/server"
	"todo-api/logger"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var (
	AppName = "TODO-API"
)

func main() {
	app := cli.NewApp()
	app.Name = "TODO-API"
	app.Usage = "Todo CLI"
	app.Version = "v1.0"

	err := logger.NewLogger(logger.Configuration{
		EnableConsole:     true,
		ConsoleJSONFormat: true,
		ConsoleLevel:      "debug",
	}, logger.InstanceLogrusLogger)
	if err != nil {
		logrus.Panic(err)
	}

	app.Commands = []*cli.Command{
		{
			Name:   "server",
			Usage:  "Starts webservice",
			Action: server.Start,
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		logrus.Panic(err)
	}
}
