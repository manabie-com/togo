package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/manabie-com/togo/pkg/database"
	"github.com/manabie-com/togo/pkg/router"
)

func main() {
	// load env
	err := godotenv.Load(".env")
	if err != nil {
		panic("load env error")
	}

	// open db and migrate
	database.Init()

	// init router then start server
	r := router.Init()
	err = r.Run()
	if err != nil {
		log.Printf("server error %s", err)
	}

	// close db
	database.Close()
}
