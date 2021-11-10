package config

import (
	"bytes"
	"encoding/json"
	"github.com/spf13/viper"
	"strings"
	"time"
)

// Config application config
type Config struct {
	Postgresql        DBConfig          `mapstructure:"DB"`
	ApplicationConfig ApplicationConfig `mapstructure:"APPLICATION"`
}

// DBConfig used to set config for database
type DBConfig struct {
	Driver string `mapstructure:"DRIVER"`
	Source string `mapstructure:"SOURCE"`
}

// loadDefaultDBConfig return default database configuration
func loadDefaultDBConfig() DBConfig {
	return DBConfig{
		Driver: "postgres",
		Source: "postgresql://mtuan:secret@localhost:5432/togo?sslmode=disable",
	}
}

// ApplicationConfig set config for task
type ApplicationConfig struct {
	TokenSecretKey string        `mapstructure:"TOKEN_SECRET_KEY"`
	TokenDuration  time.Duration `mapstructure:"TOKEN_DURATION"`
}

// loadDefaultApplicationConfig return default application config
func loadDefaultApplicationConfig() ApplicationConfig {
	return ApplicationConfig{
		TokenSecretKey: "secret",
		TokenDuration:  time.Minute * 15,
	}
}

// loadDefaultConfig return default configuration of server(for local development, ...)
func loadDefaultConfig() *Config {
	return &Config{
		Postgresql:        loadDefaultDBConfig(),
		ApplicationConfig: loadDefaultApplicationConfig(),
	}
}

// Load system env config
func Load() (*Config, error) {
	cfg := loadDefaultConfig()

	viper.SetConfigType("json")
	configBuffer, err := json.Marshal(cfg)
	if err != nil {
		return nil, err
	}

	err = viper.ReadConfig(bytes.NewBuffer(configBuffer))
	if err != nil {
		return nil, err
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()
	err = viper.Unmarshal(cfg)
	return cfg, err
}
