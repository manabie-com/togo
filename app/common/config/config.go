package config

import (
	"fmt"
	"os"

	"github.com/manabie-com/togo/app/common/adapter"

	Viper "github.com/spf13/viper"
)

const (
	ReadConfigFileErr = "read config file error"
)

// Config struct ..
type Config struct {
	Name     string            `mapstructure:"name"`
	Port     map[string]string `mapstructure:"port"`
	Version  string            `mapstructure:"version"`
	Debug    bool              `mapstructure:"debug"`
	Mongo    adapter.Mongos
	Redis    adapter.Redises
	RabbitMQ adapter.Rabbits
	Other    adapter.Others
}

var config *Config

func init() {
	var folder string

	env := os.Getenv("APPLICATION_ENV")

	switch env {
	case "master", "dev", "uat", "sandbox", "localhost":
		folder = env
	default:
		folder = "dev"
	}

	path := fmt.Sprintf("config/%v", folder)

	//Get base config
	config = new(Config)
	fetchDataToConfig(path, "base", config)

	//Get all sub config
	fetchDataToConfig(path, "mongo", &(config.Mongo))
	fetchDataToConfig(path, "rabbit", &(config.RabbitMQ))
	fetchDataToConfig(path, "other", &(config.Other))
	fetchDataToConfig(path, "redis", &(config.Redis))
}

func fetchDataToConfig(configPath, configName string, result interface{}) {
	viper := Viper.New()
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)

	err := viper.ReadInConfig() // Find and read the config file
	if err == nil {             // Handle errors reading the config file
		err = viper.Unmarshal(result)
		if err != nil { // Handle errors reading the config file
			panic(ReadConfigFileErr + err.Error())
		}
	}
}

// GetConfig func
func GetConfig() *Config {
	return config
}
