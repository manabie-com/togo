package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

//Config BE model
type Config struct {
	AppName     string
	AppVersion  string
	Environment string
	APIHost     string `toml:"api_host"`
	JWTSecret   string `toml:"jwt_secret"`

	//Configurations for DB
	DatabaseHost     string `toml:"togo_database_host"`
	DatabaseName     string `toml:"togo_database_name"`
	DatabaseUsername string `toml:"togo_database_user"`
	DatabasePort     string `toml:"togo_database_port"`
	DatabasePassword string `toml:"togo_database_password"`
	DatabaseSslMode  string `toml:"togo_database_ssl_mode"`
	DatabaseDialect  string `toml:"togo_database_dialect"`
	DatabaseTimezone string `toml:"togo_database_timezone"`
}

func (c *Config) Read() {
	var configFile string

	switch env := os.Getenv("GO_ENV"); env {
	case "staging":
		configFile = "config.staging.toml"
	case "production":
		configFile = "config.production.toml"
	case "qa":
		configFile = "config.qa.toml"
	default:
		configFile = "config.local.toml"
	}
	basePath := os.Getenv("GOPATH") + "/src/github.com/golang/manabie-com/togo/"
	configFile = basePath + configFile

	if _, err := toml.DecodeFile(configFile, &c); err != nil {
		log.Fatal(err)
	}
	log.Println("configuration environment: ", c.Environment)
}

var config = Config{}

func init() {
	log.Println("Reading configuration file")
	config.Read()
	log.Println("Read configuration file")
}

//GetConfig : export backend config
func GetConfig() *Config {
	return &config
}
