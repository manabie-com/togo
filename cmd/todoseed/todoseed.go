package main

import (
	"fmt"
	"io/ioutil"

	_ "time/tzdata"

	"github.com/manabie-com/togo/internal/pkg/config"
	p "github.com/manabie-com/togo/internal/pkg/db/postgres"
	log "github.com/sirupsen/logrus"
)

func main() {
	db, err := p.GetConnection()
	if err != nil {
		log.Fatalln(err)
	}

	// Just need master data for now
	masterSQL, err := ioutil.ReadFile(config.RootPath() + "/db/seeds/master.sql")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Start seeding master data\n")

	_, err = db.Exec(string(masterSQL))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Done!\n")
}
