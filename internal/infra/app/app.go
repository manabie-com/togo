package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/manabie-com/togo/internal/infra"

	"github.com/manabie-com/togo/internal/services/auth"

	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

type ApplicationContext struct {
	ctx     context.Context
	cfg     *infra.AppConfig
	restSrv *infra.RestService
	authSrv auth.Service
}

var ApplicationSet = wire.NewSet(
	infra.ProvideConfig,
	infra.ProvidePostgres,
	infra.ProvideRestAPIHandler,
	infra.ProvideRestService,
	infra.ProvideAuthService,
	infra.ProvideUserRepo,
	infra.ProvideUserTaskService,
	infra.ProvideEventBus,
	infra.ProvideCommandbus,
	infra.ProvideEventStore,
	infra.ProvideAggregateStore,
	infra.ProvideUserTaskCommandHandler,
	infra.ProvideUserTaskProjector,
	infra.ProvideUserConfigRepo,
	infra.ProvideUserTaskRepo,
	infra.ProvideReadRepo,
	infra.ProvidePostgresSlave,
	infra.ProvideTaskRepo,
	infra.ProvideTaskHandler,
)

func (a *ApplicationContext) Commands() *cli.App {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		a.Serve(),
	}

	return app
}

// HandleSigterm -- Handles Ctrl+C or most other means of "controlled" shutdown gracefully.
// Invokes the supplied func before exiting.
func HandleSigterm(handleExit func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTRAP)
	go func() {
		<-c
		logrus.Infof("Handler shutdown signal in main process...")
		handleExit()
		os.Exit(1)
	}()
}
