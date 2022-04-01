package config

import (
	"github.com/caarlos0/env/v5"
	"github.com/joho/godotenv"
)

// Configuration represents the server configuration
type Configuration struct {
	Host         string `env:"HOST"`
	Port         int    `env:"PORT"`
	ReadTimeout  int    `env:"READ_TIMEOUT"`
	WriteTimeout int    `env:"WRITE_TIMEOUT"`
}

// Load returns Configuration struct
func Load() (*Configuration, error) {
	basePath := ""
	if err := godotenv.Load(basePath + ".env"); err != nil {
		return nil, err
	}

	cfg := new(Configuration)
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
