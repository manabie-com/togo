package configs

import (
	"log"

	"github.com/spf13/viper"
)

var C Config

func ReadConfig() {
	config := &C
	configFile := "configs/config.yaml"

	viper.SetConfigFile(configFile)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("☠️ cannot read configuration", err)
	}
	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("☠️ environment can't be loaded: ", err)
	}
}
