package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config contain config value
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	DBSourceTest  string `mapstructure:"DB_SOURCE_TEST"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig load config file
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("main")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	// Find and read the config file
	err = viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	err = viper.Unmarshal(&config)
	return
}
