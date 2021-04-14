package infra

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/env"
)

type PostgresConfig struct {
	Address      string `json:"addr"`
	SlaveAddress string `json:"slaveaddr"`
	Database     string `json:"db"`
	Username     string `json:"user"`
	Password     string `json:"password"`
	ReadTimeout  int    `json:"readtimeout"`
}

type AppConfig struct {
	Port      int            `json:"port"`
	Postgres  PostgresConfig `json:"postgres"`
	SecretJWT string         `json:"secretjwt"`
}

func ProvideConfig() (*AppConfig, error) {
	cfg := &AppConfig{}
	conf, err := config.NewConfig(config.WithSource(env.NewSource()))
	if err != nil {
		return nil, err
	}

	if err := conf.Scan(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
