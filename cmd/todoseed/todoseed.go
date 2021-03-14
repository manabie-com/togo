package main

import (
	"fmt"
	"io/ioutil"

	_ "time/tzdata"

	p "github.com/manabie-com/togo/internal/pkg/db/postgres"
	"github.com/manabie-com/togo/internal/pkg/projectpath"
	log "github.com/sirupsen/logrus"
)

func main() {
	db, err := p.GetConnection()
	if err != nil {
		log.Fatalln(err)
	}

	// Just need master data for now
	masterSQL, err := ioutil.ReadFile(projectpath.RootPath() + "/db/seeds/master.sql")
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
