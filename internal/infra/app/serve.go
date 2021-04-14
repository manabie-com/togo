package app

import (
	"github.com/urfave/cli"
)

// Serve start application
func (a *ApplicationContext) Serve() cli.Command {
	return cli.Command{
		Name:  "serve",
		Usage: "Start Service",
		Action: func(c *cli.Context) error {
			a.restSrv.MustStart()
			return nil
		},
	}
}
