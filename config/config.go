package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"

	"github.com/khangjig/togo/client/logger"
	"github.com/khangjig/togo/codetype"
)

var config *Config

type Config struct {
	Port           string             `envconfig:"PORT"`
	IsDebug        bool               `envconfig:"IS_DEBUG"`
	Stage          codetype.StageType `envconfig:"STAGE"`
	TokenSecretKey string             `envconfig:"TOKEN_SECRET_KEY"`

	MySQL struct {
		Host   string `envconfig:"DB_HOST"`
		Port   string `envconfig:"DB_PORT"`
		User   string `envconfig:"DB_USER"`
		Pass   string `envconfig:"DB_PASS"`
		DBName string `envconfig:"DB_NAME"`
	}

	Redis struct {
		Host    string `envconfig:"REDIS_HOST"`
		Port    string `envconfig:"REDIS_PORT"`
		DB      int    `envconfig:"REDIS_DB"`
		User    string `envconfig:"REDIS_USER"`
		Pass    string `envconfig:"REDIS_PASS"`
		Timeout int    `envconfig:"REDIS_TIMEOUT"`
	}

	HealthCheck struct {
		EndPoint string `envconfig:"HEALTH_CHECK_ENDPOINT"`
	}
}

func init() {
	_ = godotenv.Load()
	config = &Config{}

	err := envconfig.Process("", config)
	if err != nil {
		logger.GetLogger().Fatal(errors.Wrap(err, "Failed to decode config env").Error())
	}

	config.Stage.UpCase()
}

func GetConfig() *Config {
	return config
}
