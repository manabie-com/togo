package main

import (
	"togo/globals/database"
	"togo/migration"
	"togo/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	database.InitDBConnection()
	migration.Migrate(database.SQL)

	routes := routes.InitRoutes()

	routes.Run(":8080")
}