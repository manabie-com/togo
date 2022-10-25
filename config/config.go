package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
)

type Config struct {
	Port      string
	TimeOut   string
	DBPort    string
	DBHost    string
	DBUser    string
	DBPass    string
	DBName    string
	DBSsl     string
	LimitTask int
}

var config Config

func NewFromFile() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file", err)
		return nil, err
	}
	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("unable to decode into struct", err)
	}
	return &config, nil
}

func New() (*Config, error) {
	godotenv.Load()
	config.Port = os.Getenv("PORT")
	config.TimeOut = os.Getenv("TIME_OUT")
	config.DBHost = os.Getenv("DB_HOST")
	config.DBPort = os.Getenv("DB_PORT")
	config.DBUser = os.Getenv("DB_USER")
	config.DBPass = os.Getenv("DB_PASS")
	config.DBName = os.Getenv("DB_NAME")
	config.DBSsl = os.Getenv("DB_SSL")

	if config.Port == "" {
		config.Port = "8080"
	}
	if config.TimeOut == "" {
		config.TimeOut = "30"
	}
	if config.DBSsl == "" {
		config.DBSsl = "disable"
	}

	config.LimitTask, _ = strconv.Atoi(os.Getenv("LIMIT_TASK"))
	if config.LimitTask == 0 {
		// set default limit task is 10
		config.LimitTask = 10
	}

	return &config, nil
}
