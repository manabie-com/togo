package config

import (
	"github.com/spf13/viper"
)

func LoadConfig(path string, configFileName string, output interface{}) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(configFileName)
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(&output)
	return
}
