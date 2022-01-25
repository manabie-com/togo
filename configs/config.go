package configs

import (
	"github.com/trinhdaiphuc/env_config"
	"time"
)

type Config struct {
	Port        int           `env:"PORT,8080"`
	IdleTimeout time.Duration `env:"IDLE_TIMEOUT,5s"`
	JwtSecret   string        `env:"JWT_SECRET,secret"`
	DBHost      string        `env:"DB_HOST,localhost"`
	DBPort      int           `env:"DB_PORT,3306"`
	DBUser      string        `env:"DB_USER,root"`
	DBPassword  string        `env:"DB_PASSWORD,root"`
	DBName      string        `env:"DB_NAME,togo"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env_config.EnvStruct(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
