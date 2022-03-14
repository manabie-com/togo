package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/triet-truong/todo/utils"
)

var appConfig config

type config struct {
	appPort int

	dbHost     string
	dbPort     int
	dbUsername string
	dbPassword string
	dbName     string
	cacheHost  string
	cachePort  int
}

// Load load config
func Load() {
	env := os.Getenv("ENVIRONMENT")
	if env == "LOCAL" {
		fmt.Println(env)
		viper.SetConfigFile(".env")
		err := viper.ReadInConfig()
		utils.FatalLog(err)
	}

	viper.AutomaticEnv()
	appConfig = config{
		appPort:    viper.GetInt("PORT"),
		dbHost:     viper.GetString("DB_HOST"),
		dbPort:     viper.GetInt("DB_PORT"),
		dbUsername: viper.GetString("DB_USER"),
		dbPassword: viper.GetString("DB_PASS"),
		dbName:     viper.GetString("DB_NAME"),
		cacheHost:  viper.GetString("CACHE_HOST"),
		cachePort:  viper.GetInt("CACHE_PORT"),
	}
}

func DatabaseDSN() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?allowNativePasswords=True&parseTime=True", appConfig.dbUsername, appConfig.dbPassword, appConfig.dbHost, appConfig.dbPort, appConfig.dbName)
}

func CacheConnectioURL() string {
	return fmt.Sprintf("%v:%v", appConfig.cacheHost, appConfig.cachePort)
}

// Port return server port
func Port() int {
	return appConfig.appPort
}
