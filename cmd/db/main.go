package main

import (
	"github.com/manabie-com/togo/config"
)

func main() {
	config.LoadEnv("")
	config.ConnectDB()
	config.Migrate()
	config.Seed()
}
