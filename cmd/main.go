package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/manabie-com/togo/internal/core/port"
	"github.com/manabie-com/togo/internal/core/service"
	"github.com/manabie-com/togo/internal/core/validator"
	"github.com/manabie-com/togo/internal/handler"
	"github.com/manabie-com/togo/internal/repository"
	"github.com/manabie-com/togo/pkg/database"

	"go.uber.org/dig"
)

const (
	gJwtKey = "wqGyEBBfPK9w3Lxw"
)

var (
	exitServer context.CancelFunc
)

func main() {
	container := buildContainer()
	container.Invoke(func(ctx context.Context, serverHandler handler.HttpHandler) {
		// Start server
		fmt.Println("starting Gateway Api server")

		serverExited := make(chan bool)
		go func() {
			err := serverHandler.Begin(ctx, ":5050")
			if err != nil {
				fmt.Println("start Gateway Api server failed", err.Error())
			}
			close(serverExited)
		}()

		sig := make(chan os.Signal)
		signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
		<-sig
		fmt.Println("received interrupt. Exit")

		// Stop server
		exitServer()
		<-serverExited

		fmt.Println("stopped Gateway Api server")
	})
}

func buildContainer() *dig.Container {
	container := dig.New()

	// Setup context
	container.Provide(func() context.Context {
		var ctx context.Context
		ctx, exitServer = context.WithCancel(context.Background())
		return ctx
	})

	// Setup database
	container.Provide(func() database.Database {
		db := database.NewDatabase(database.NewSqliteConnector("./data.db"))
		err := db.Connect(nil)
		if err != nil {
			panic("setup Flow repository failed: " + err.Error())
		}
		return db
	})

	// Setup Task repository
	container.Provide(repository.NewTaskRepository)

	// Setup Task validator
	container.Provide(validator.NewTaskValidator)

	// Setup Task service
	container.Provide(service.NewTaskService)

	// Setup JWT service
	container.Provide(func() port.JwtService {
		return service.NewJwtService(gJwtKey)
	})

	// Setup handler
	container.Provide(handler.NewHttpHandler)

	return container
}
