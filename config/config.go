package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DbUsername string `required:"true" split_words:"true"`
	DbPassword string `required:"true" split_words:"true"`
	DbName     string `required:"true" split_words:"true"`
	DbHost     string `required:"true" split_words:"true"`
	DbPort     int    `required:"true" split_words:"true"`
	SslMode    string `required:"true" split_words:"true"`
	Port       int    `required:"true" split_words:"true"`
}

func New() (*Config, error) {
	var c Config

	err := envconfig.Process("", &c)

	if err != nil {
		return nil, err
	}

	return &c, nil
}
