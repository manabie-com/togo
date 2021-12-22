package server

import (
	"time"
	"todo-api/config"
	"todo-api/logger"
	"todo-api/migrations"
	"todo-api/service"
	"todo-api/src/router"

	"github.com/labstack/echo/v4"
	"github.com/urfave/cli/v2"
	"gopkg.in/tylerb/graceful.v1"
)

func Start(c *cli.Context) error {
	app := echo.New()

	cfg, err := config.New()
	if err != nil {
		return err
	}

	log := logger.GetLogger()
	log.Infof("Application %s", c.App.Name)
	db := config.Database()
	if !migrations.MigrateUp(config.GetDBDriver(), config.GetDBUrl()) {
		log.Errorf("Migration is failed")
	}

	s := &service.Service{
		DB:     db,
		Logger: log,
		Config: cfg,
	}

	router.Init(app, s)

	app.Server.Addr = ":" + cfg.Port
	log.Infof("Server started at %s", app.Server.Addr)
	graceful.ListenAndServe(app.Server, 5*time.Second)
	return nil
}
