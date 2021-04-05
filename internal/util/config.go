package util

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

var Conf *Config

type Config struct {
	PostgresDriver   string `mapstructure:"POSTGRES_DRIVER"`
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
	PostgresDB       string `mapstructure:"POSTGRES_DB"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresSLLMode  string `mapstructure:"POSTGRES_SLLMODE"`

	SecretKey  string        `mapstructure:"SECRET_KEY"`
	Timeout    time.Duration `mapstructure:"TIMEOUT"`
	FormatDate string        `mapstructure:"FORMAT_DATE"`
	Address    string        `mapstructure:"Address"`
	DBType     string        `mapstructure:"DB_TYPE"`

	SqlLiteDriver string `mapstructure:"SQL_LITE_DRIVER"`
	SqlLiteFile   string `mapstructure:"SQL_LITE_FILE"`
}

func LoadConfig(path string) error {
	viper.AddConfigPath(path)
	stage := os.Getenv("ENVIRONMENT")
	name := "app." + stage
	viper.SetConfigName(name)
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&Conf)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) ConnectionString() string {
	if c.DBType == c.SqlLiteDriver {
		return c.SqlLiteFile
	}
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.PostgresHost, c.PostgresPort, c.PostgresUser, c.PostgresPassword, c.PostgresDB, c.PostgresSLLMode)
}
