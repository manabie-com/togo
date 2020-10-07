package app

import (
	"fmt"
	"log"

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
		fmt.Println("App running...")
		// TODO:
		db, err := util.CreateConnectionDB()
		if err != nil {
			log.Fatalf("Connect DB error: %v", err)
		}
		e := initEcho()

		// Load modules
		module.LoadModules(e, db)

		// Start server
		port := config.Cfg.Port
		if config.Cfg.Env == "local" {
			log.Fatal(e.Start("localhost:" + port))
		} else {
			log.Fatal(e.Start(":" + port))
		}
	}
}
