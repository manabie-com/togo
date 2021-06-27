package config

import (
	"github.com/kelseyhightower/envconfig"

	"github.com/manabie-com/togo/internal/pkgs/clients"
)


type Config struct {
	DB clients.PSQLConfig
	HTTP HTTPConf

}

type HTTPConf struct {
	Addr string `envconfig:"HTTP_ADDR" default:"0.0.0.0:5050"`
}

func Load() (*Config, error) {
	c := Config{}
	err := envconfig.Process("", &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
