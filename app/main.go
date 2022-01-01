package main

import (
	"os"

	httpServer "github.com/manabie-com/togo/app/web/http_server"

	"github.com/manabie-com/togo/app/common/config"
	"github.com/urfave/cli"
)

var cfg = config.GetConfig()

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		// ##### Web api #####
		{
			Name:    "api",
			Aliases: []string{"s"},
			Action: func(c *cli.Context) {
				httpServer.Run()
			},
		},
	}

	if len(os.Args) == 1 {
		os.Args = append(os.Args, "s")
	}

	app.Run(os.Args)
}
