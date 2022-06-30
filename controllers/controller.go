package controllers

import (
	"context"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"sync"
	"time"
	"togo-thdung002/config"
	"togo-thdung002/entities/sqlite"
)

type Controller struct {
	e *echo.Echo
	sync.Mutex
	conf *config.Config
	db   *sqlite.LiteDB
}

func NewController(conf *config.Config, sqlitedb *gorm.DB) *Controller {
	db := sqlite.NewDB(sqlitedb)

	ctrl := Controller{
		db:   db,
		conf: conf,
	}
	ctrl.loadMux()
	log.Infof("INF: Loading API Listener on %s\n", ":8800")

	return &ctrl
}
func (ctrl *Controller) Load() error {
	ctrl.Lock()
	defer ctrl.Unlock()
	return nil
}

// Start is non-blocking
func (ctrl *Controller) Start() {
	ctrl.Lock()
	defer ctrl.Unlock()
}

// Stop is non-blocking
func (ctrl *Controller) Stop() {
	ctrl.Lock()
	defer ctrl.Unlock()
}

func (ctrl *Controller) ListenAndServe() error {

	apiServer := &http.Server{
		Addr:         ctrl.conf.API.Listen,
		ErrorLog:     nil,
		ReadTimeout:  5 * time.Minute,
		WriteTimeout: 5 * time.Minute,
	}
	ctrl.e.HideBanner = false
	ctrl.e.Debug = false
	return ctrl.e.StartServer(apiServer)
}

func (ctrl *Controller) Shutdown(ctx context.Context) error {
	ctrl.e.Shutdown(ctx)
	return nil
}
