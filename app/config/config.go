package config

import (
	"ansidev.xyz/pkg/db"
	"ansidev.xyz/pkg/rds"
)

var (
	AppConfig Config
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	LogLevel        string `mapstructure:"LOG_LEVEL"`
	Host            string `mapstructure:"HOST"`
	Port            int    `mapstructure:"PORT"`
	db.SqlDbConfig  `mapstructure:",squash"`
	rds.RedisConfig `mapstructure:",squash"`
	TokenTTL        int `mapstructure:"TOKEN_TTL"`
}
