package utils

import "github.com/spf13/viper"

type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	viper.Unmarshal(&config)
	return
}
