package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/manabie-com/togo/internal/logs"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/manabie-com/togo/internal/transport"
	"github.com/manabie-com/togo/internal/util"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	logger := logs.WithPrefix("main")
	_, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		logger.Panic("error opening db", "process", err.Error())
	}

	err = util.LoadConfig("./configs")
	if err != nil {
		logger.Panic("error loading config", "process", err.Error())
	}

	// serving and return error
	postgres := postgres.NewPostgres()
	server := transport.NewServer(postgres)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt)
	go func() {
		if err := server.Start(util.Conf.Address); err != nil {
			logger.Info("Cannot start server", "process", err)
			return
		}
	}()

	logger.Info("Server is running at", "process", nil)

	fmt.Println(quit)
	a := <-quit
	fmt.Println(a)
	defer func() {
		fmt.Println("2")

	}()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println(err)
	}

	fmt.Println("server is stop")
}
