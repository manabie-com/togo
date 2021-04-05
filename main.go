package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/manabie-com/togo/internal/logs"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/postgres"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	"github.com/manabie-com/togo/internal/transport"
	"github.com/manabie-com/togo/internal/util"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	logger := logs.WithPrefix("main")
	err := util.LoadConfig("./configs")
	if err != nil {
		logger.Panic("error loading config", err.Error())
	}

	// serving
	store := initStore()
	server := transport.NewServer(store)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := server.Start(util.Conf.Address); err != nil {
			logger.Error("Cannot start server", err.Error())
			return
		}
	}()

	logger.Info(fmt.Sprintf("Server is running at %v", util.Conf.Address), nil)

	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println(err)
	}

	fmt.Println("server is stop")
}

func initStore() storages.Store {
	if util.Conf.DBType == util.Conf.SqlLiteDriver {
		return sqllite.NewLitDB()
	}

	return postgres.NewPostgres()
}
