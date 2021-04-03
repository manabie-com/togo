package util

import (
	"os"

	"github.com/spf13/viper"
)

var Conf *Config

type Config struct {
	PostgresDriver   string `mapstructure:"POSGRES_DRIVER"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
	PostgresDB       string `mapstructure:"POSTGRES_DB"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresSLLMode  string `mapstructure:"POSTGRES_SLLMODE"`

	SecretKey string `mapstructure:"SECRET_KEY"`
}

func LoadConfig(path string) error {
	viper.AddConfigPath(path)
	a := os.Getenv("ENVIRONMENT")
	name := "app." + a
	viper.SetConfigName(name)
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&Conf)
	if err != nil {
		return err
	}

	return nil
}
