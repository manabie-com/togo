package configs

import (
	"togo/pkg/logger"

	"github.com/spf13/viper"
)

var C Config

func ReadConfig() {
	config := &C
	configFile := "configs/config.yaml"

	viper.SetConfigFile(configFile)

	err := viper.ReadInConfig()
	if err != nil {
		logger.L.Sugar().Fatalf("☠️ cannot read configuration", err)
	}
	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	if err != nil {
		logger.L.Sugar().Fatalf("☠️ environment can't be loaded: ", err)
	}
}
