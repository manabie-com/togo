package main

import (
	"fmt"
	"log"

	"example.com/m/v2/cmd/middlewares"
	"example.com/m/v2/constants"
	"example.com/m/v2/internal/api/handlers"
	"example.com/m/v2/internal/api/routes"
	"example.com/m/v2/internal/driver"
	"example.com/m/v2/utils"

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
		r.Use(gin.Logger())
		// Recovery middleware recovers from any panics and writes a 500 if there was one.
		r.Use(gin.Recovery())
		r.Use(middlewares.SetDefaultMiddleWare())
		r.Use(middlewares.ValidateToken())
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
