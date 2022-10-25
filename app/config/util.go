package config

import (
	"ansidev.xyz/pkg/log"
	"github.com/spf13/viper"
)

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string, configFile string, config *Config) {
	viper.AddConfigPath(path)
	viper.SetConfigName(configFile)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Debug("Could not load config:", err)
		panic(err.Error())
	}

	err = viper.Unmarshal(config)

	if err != nil {
		log.Fatal("Could not parse config:", err)
		panic(err.Error())
	}
}
