package config

import (
	"os"
)
// Config Global Variable
var Config *Conf

type Conf struct {
	AppEnv	string
	AppDebug bool
	AppUrl  string
	AppPort int
	DBHost 	string
	DBPort 	string   
	DBUser 	string
	DBPass 	string
	DBName 	string
}

func InitConfig() {
	// TODO: Load config from .env file
	config := &Conf{
		AppEnv: "local",
		AppDebug: false,
		AppPort: 5050,
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBUser: os.Getenv("DB_USERNAME"),
		DBPass: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
	}
	Config = config
}