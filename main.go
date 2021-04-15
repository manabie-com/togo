package main

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/router"
	"github.com/manabie-com/togo/internal/util"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	configPath := os.Getenv("TOGO_CONFIG_PATH")
	if configPath == "" {
		configPath = "./"
	}

	conf, err := config.InitializeConfig(configPath)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to read configuration file")
		return
	}
	db, err := sql.Open(conf.DB.DriverName, util.GetConnectionString(conf.DB))
	if err != nil {
		log.Fatal().Err(err).Msg("unable to open DB")
		return
	}
	defer db.Close()

	handler := router.NewRouter(db, conf.DB.DriverName, conf.JWTKey)
	server := &http.Server{
		Addr:         ":5050",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Info().Msg("Togo is ready")
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal().Err(err).Msg("unable to start the service")
		}
	}()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	<-signals

	ctx, cancelFn := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelFn()
	if err := server.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("unable to shutdown the service")
	}
}
