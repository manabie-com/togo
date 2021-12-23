package utils

import (
	"github.com/spf13/viper"
)

type Config struct {
	MYSQL_HOST         string `mapstructure:"MYSQL_HOST"`
	MYSQL_USERNAME     string `mapstructure:"MYSQL_USERNAME"`
	MYSQL_PASSWORD     string `mapstructure:"MYSQL_PASSWORD"`
	MYSQL_DB           string `mapstructure:"MYSQL_DB"`
	LIMIT_TASK_PER_DAY string `mapstructure:"LIMIT_TASK_PER_DAY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app") //our file is app.env
	viper.SetConfigType("env") //our file is app.env

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
