package config

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Port     int
	MySqlUri string
}

var configuration Configuration
var defaultconfig Configuration

func GetConfig() Configuration {
	if configuration == defaultconfig {
		fmt.Println("get config")
		if err := gonfig.GetConf("config.json", &configuration); err != nil {
			fmt.Println(err)
			return Configuration{
				Port:     5000,
				MySqlUri: "user_testdb:password@tcp(test-db:3306)/testdb",
			}
		}
		if configuration == defaultconfig {
			return Configuration{
				Port:     5000,
				MySqlUri: "user_testdb:password@tcp(test-db:3306)/testdb",
			}
		}
	}
	return configuration
}
