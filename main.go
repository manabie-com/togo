package main

import (
	"context"
	"flag"
	"github.com/manabie-com/togo/config"
	"github.com/manabie-com/togo/internal/api"
	"github.com/manabie-com/togo/internal/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var mbLogger logger.Logger

func main() {
	state := flag.String("state", "local", "state of service")
	mbLogger = logger.WithPrefix("main")
	cfg, err := config.Load(state)
	if err != nil {
		mbLogger.Panicln(err)
	}
	var server *http.Server
	go func() {
		server = initRestfulAPI(cfg)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			mbLogger.Panicf("Fail to listen and server: %v", err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	<-signals

	mbLogger.Info("shutting server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		mbLogger.Errorf("Fail to listen and server: %v", err)
	}
	mbLogger.Info("shutdown server")
}

func initRestfulAPI(cfg *config.Config) *http.Server {
	mbLogger.Info("Start server")
	mbLogger.Infof("%s:%s", cfg.RestfulAPI.Host, cfg.RestfulAPI.Port)
	server, err := api.CreateAPIEngine(cfg)
	if err != nil {
		mbLogger.Panicf("Fail init server: %v", err)
		return nil
	}
	return server
}
