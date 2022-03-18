package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/manabie-com/togo/server"

	"github.com/manabie-com/togo/registry"

	"github.com/manabie-com/togo/core/config"
)

func main() {
	bytes, err := config.Asset("config.json")
	if err != nil {
		log.Fatal(err)
	}
	var cfg config.Config
	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	r, err := registry.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	db, err := r.DB.ManabieDB.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	testDB, err := r.DB.TestManabieDB.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer testDB.Close()

	if err = r.DB.Migrate(); err != nil {
		log.Fatal(err)
	}
	if err = r.DB.MigrateTestDB(); err != nil {
		log.Fatal(err)
	}
	errs := make(chan error, 1)
	go func() {
		fmt.Println("HTTP server content management listening on 8080")
		errs <- http.ListenAndServe(fmt.Sprintf(":%s", "8080"), server.New(r))
	}()
	log.Println(fmt.Sprintf("exiting (%v)", <-errs))
}
