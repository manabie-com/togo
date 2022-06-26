package main

import (
	"fmt"
	"log"

	"example.com/m/v2/constants"
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

	log.Fatal(r.Run(fmt.Sprintf(":%s", utils.Env.Port)))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
