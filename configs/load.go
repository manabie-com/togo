package configs

import (
	"togo/pkg/logger"

	"github.com/spf13/viper"
)

func ReadConfig() *Config {
	config := &Config{}
	configFile := "configs/config.yaml"

	viper.SetConfigFile(configFile)

	err := viper.ReadInConfig()
	if err != nil {
		logger.L.Sugar().Fatalf("☠️ cannot read configuration at path: %s with err : %v", configFile, err)
	}
	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	if err != nil {
		logger.L.Sugar().Fatalf("☠️ environment can't be loaded: ", err)
	}
	return config
}
