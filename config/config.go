package config

import (
	"github.com/caarlos0/env/v5"
	"github.com/joho/godotenv"
)

// Configuration represents the server configuration
type Configuration struct {
	// Server configuration parameters
	ServerPort         int  `env:"SERVER_PORT"`
	ServerReadTimeout  int  `env:"SERVER_READ_TIMEOUT"`
	ServerWriteTimeout int  `env:"SERVER_WRITE_TIMEOUT"`
	ServerDebug        bool `env:"SERVER_DEBUG"`

	// Database configuration parameters
	DbHost     string `env:"POSTGRES_HOST"`
	DbPort     int    `env:"POSTGRES_PORT"`
	DbUser     string `env:"POSTGRES_USER"`
	DbPassword string `env:"POSTGRES_PASSWORD"`
	DbName     string `env:"POSTGRES_DB"`
	DbSslMode  string `env:"POSTGRES_SSL_MODE"`
	DbLog      bool   `env:"POSTGRES_LOG"`

	// JWT configuration parameters
	JwtSecret    string `env:"JWT_SECRET"`
	JwtDuration  int    `env:"JWT_DURATION"`
	JwtAlgorithm string `env:"JWT_ALGORITHM"`
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
