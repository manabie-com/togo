package main

import (
	"fmt"
	"os"
	"strconv"

	_ "time/tzdata"

	"github.com/golang-migrate/migrate/v4"
	"github.com/manabie-com/togo/internal/pkg/config"
	p "github.com/manabie-com/togo/internal/pkg/db/postgres"
	log "github.com/sirupsen/logrus"
)

func main() {
	migratePath := config.RootPath() + "/db/migrations"
	m, err := p.GetMigrateInstance(migratePath)
	if err != nil {
		log.Fatalln(err)
	}

	defer m.Close()

	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatalln("Migration command either 'up' or 'down' or 'force' is required")
	}
	if len(args) > 2 {
		log.Fatalln("Invalid command arguments")
	}

	switch args[0] {
	case "up":
		if len(args) == 2 && args[1] == "all" {
			err = m.Up()
		} else {
			err = m.Steps(1)
		}
	case "down":
		if len(args) == 2 && args[1] == "all" {
			err = m.Down()
		} else {
			err = m.Steps(-1)
		}
	case "force":
		if len(args) < 2 {
			log.Fatalln("Force require a version number. For ex: force 3")
		}

		version, parseErr := strconv.Atoi(args[1])
		if parseErr != nil {
			log.Fatalf("Invalid version: %v\n", parseErr)
		}

		err = m.Force(version)
	default:
		log.Fatalln("Invalid migration command")

	}

	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("%+v\n", err)
	}

	if err == migrate.ErrNoChange {
		fmt.Println("No new changes!")
	}

	fmt.Println("Done running migration!")
}
