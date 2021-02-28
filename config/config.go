package config

import (
	cc "github.com/manabie-com/togo/pkg/common/config"
)

type Config struct {
	cc.Postgres `yaml:"postgres"`
	JWTSecret   string  `yaml:"jwt_secret"`
	MaxTodo     int     `yaml:"max_todo"`
	HTTP        cc.HTTP `yaml:"http"`
}

func Default() Config {
	return Config{
		Postgres:  cc.DefaultPostgres(),
		JWTSecret: "secret",
		MaxTodo:   5,
		HTTP: cc.HTTP{
			Host: "",
			Port: 5050,
		},
	}
}

func Load() (cfg Config, err error) {
	err = cc.LoadWithDefault(&cfg, Default())
	return cfg, err
}
