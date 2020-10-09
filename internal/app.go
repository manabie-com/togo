package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/module"
	"github.com/manabie-com/togo/internal/util"
)

func initEcho() *echo.Echo {
	e := echo.New()
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_unix} ${id} ${method} ${uri} ${status} ${latency_human}\n",
	}))

	e.Use(middleware.CORS())

	return e
}

// Run func
func Run() {

	//Load env
	if err := config.Load(); err != nil {
		log.Fatalf("Load config error: %v", err)
	} else {
		fmt.Println("App started...")

		db, err := util.CreateConnectionDB()
		defer func() {
			fmt.Println("App shutdown")
			dbSQL, ok := db.DB()
			if ok != nil {
				defer dbSQL.Close()
			}
		}()

		if err != nil {
			log.Fatalf("Connect DB error: %v", err)
		}
		e := initEcho()

		// Load modules
		module.LoadModules(e, db)

		//TODO: https://echo.labstack.com/cookbook/graceful-shutdown

		// Start server
		go func() {
			address := ":" + config.Cfg.Port
			if config.Cfg.Env == "local" {
				address = "localhost:" + config.Cfg.Port
			}
			fmt.Println(e.Start(address))
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
}
