package main

import (
	"fmt"
	"log"

	"github.com/manabie-com/togo/cmd/middlewares"
	"github.com/manabie-com/togo/constants"
	"github.com/manabie-com/togo/internal/api/handlers"
	"github.com/manabie-com/togo/internal/api/routes"
	"github.com/manabie-com/togo/internal/driver"
	"github.com/manabie-com/togo/utils"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func run() error {
	// Creates a new Gin instance.
	r := gin.Default()

	// Load File Environment
	utils.LoadEnv(constants.FileEnvironment)

	// Connect to Database
	db, err := driver.ConnectDatabase()
	if err != nil {
		return errors.Wrap(err, "Fail to ConnectDatabase")
	}

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	{ // Set Middlewares
		// Midlewares
		// Recovery middleware recovers from any panics and writes a 500 if there was one.
		r.Use(gin.Recovery())
		r.Use(middlewares.SetDefaultMiddleWare())
	}

	{ //Set Repository
		usecases := handlers.NewUseCase(db)
		handlers.SetMainUseCase(&usecases)
		//Set Route
		routes.SetupRoute(r, usecases)
	}

	log.Fatal(r.Run(fmt.Sprintf(":%s", utils.Env.Port)))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
