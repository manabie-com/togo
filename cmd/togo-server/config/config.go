package config

import (
	cc "github.com/manabie-com/togo/pkg/common/config"
)

type Config struct {
	cc.Postgres `yaml:"postgres"`
	JWTSecret   string `yaml:"jwt_secret"`
}

func Default() Config {
	return Config{
		Postgres:  cc.DefaultPostgres(),
		JWTSecret: "secret",
	}
}

func Load() (cfg Config, err error) {
	err = cc.LoadWithDefault(&cfg, Default())
	return cfg, err
}