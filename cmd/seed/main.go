package main

import (
	"github.com/manabie-com/togo/config"
	"github.com/manabie-com/togo/db"
)

func main() {
	config.LoadEnv()

	conn := db.ConnectDB(false)

	defer db.DisconnectDB(conn)

	db.Seed(conn)
}
