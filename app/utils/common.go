package utils

import (
	"github.com/huuthuan-nguyen/manabie/config"
	"log"
)

// ReadConfig /**
func ReadConfig() *config.Config {
	c := &config.Config{}
	// get the config path from command arguments
	configPath, err := config.ParseConfigFlags()
	if err != nil {
		log.Printf("Config File path is can not be loaded: %v\n", err)
	}

	// parse config from yml
	err = config.ReadENVFile(c, configPath)
	if err != nil {
		log.Fatalf("Fail to load config from env: %v", err)
	}

	return c
}
