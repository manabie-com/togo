// Package main
package main

import (
	"log"

	"github.com/manabie-com/togo/initializer"
)

func main() {
	if err := initializer.InitApplication(".env"); err != nil {
		log.Panic("Failed while Init application: err=", err)
	}

	initializer.GlobalConfig.GinRouter().Run(":8080")
}
