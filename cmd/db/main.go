package main

import (
	"github.com/manabie-com/togo/config"
	"github.com/manabie-com/togo/db"
)

func main() {
	config.LoadEnv("")
	db.ConnectDB()
	db.Migrate()
	db.Seed()
}
