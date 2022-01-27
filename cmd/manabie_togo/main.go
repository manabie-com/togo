package main

import (
	"encoding/json"
	"log"

	"github.com/manabie-com/togo/core/registry"

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
	if err = r.DB.Migrate(); err != nil {
		log.Fatal(err)
	}
}
