package config

import (
	"log"

	"github.com/pkg/errors"
	"github.com/kelseyhightower/envconfig"
	"github.com/joho/godotenv"
)

type config struct {
	JWTKey           string `envconfig:"JWTKey" default:""`
	HttpInterface    string `envconfig:"HTTP_INTERFACE" default:"root"`
	HttpPort         int    `envconfig:"HTTP_PORT" default:"5050"`
	PostgresHost     string `envconfig:"POSTGRES_HOST" default:"localhost"`
	PostgresPort     int    `envconfig:"POSTGRES_PORT" default:"5432"`
	PostgresUser     string `envconfig:"POSTGRES_USER" default:"postgres"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD" default:"123456"`
	PostgresDb       string `envconfig:"POSTGRES_DB" default:"test"`
	RedisHost        string `envconfig:"REDIS_HOST" default:"localhost"`
	RedisPort        int    `envconfig:"REDIS_PORT" default:"6379"`
	RedisPassword    string `envconfig:"REDIS_PASSWORD" default:""`
	RedisDb          int    `envconfig:"REDIS_DB" default:"1"`
}

var Values config

func Setup() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := envconfig.Process("", &Values); err != nil {
		err = errors.Wrap(err, "parse environment variables")
		return err
	}

	return nil
}
