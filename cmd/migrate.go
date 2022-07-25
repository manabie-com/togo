package main

import (
	"flag"
	"github.com/huuthuan-nguyen/manabie/app/utils"
	"github.com/huuthuan-nguyen/manabie/migration"
	"log"
)

var rollback = flag.Bool("rollback", false, "Rollback to previous migration.")
var reset = flag.Bool("reset", false, "Reset the database and run all migrations.")

func main() {
	flag.Parse()
	var err error

	config := utils.ReadConfig()
	engine := migration.NewEngine(config)
	// if rollback
	if *rollback {
		err = engine.Rollback()
	} else if *reset {
		err = engine.Reset()
	} else { // default is migrate
		err = engine.Migrate()
	}

	if err != nil {
		log.Fatalf("Error migrating database:%v\n", err)
	}
}
