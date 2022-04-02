package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

// Config represents server specific config
type Config struct {
	Port         int
	ReadTimeout  int
	WriteTimeout int
	Debug        bool
}

// New creates new echo server with customizable configuration
func New(cfg *Config) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger(), middleware.Recover())
	e.GET("/", healthCheck)
	e.Validator = newValidator()
	e.HTTPErrorHandler = newErrorHandler(e).HandlerFunc
	e.Binder = newBinder()
	e.Server.Addr = fmt.Sprintf(":%d", cfg.Port)
	e.Server.ReadTimeout = time.Duration(cfg.ReadTimeout) * time.Second
	e.Server.WriteTimeout = time.Duration(cfg.WriteTimeout) * time.Second
	e.Debug = cfg.Debug
	if e.Debug {
		e.Logger.SetLevel(log.DEBUG)
	} else {
		e.Logger.SetLevel(log.ERROR)
	}

	return e
}

// Start starts echo server
func Start(e *echo.Echo) {
	// Start server
	go func() {
		if err := e.StartServer(e.Server); err != nil {
			if err == http.ErrServerClosed {
				e.Logger.Info("Shutting down the server")
			} else {
				e.Logger.Errorf("Error shutting down the server: ", err)
			}
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
