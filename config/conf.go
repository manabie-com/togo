package config

import (
	"fmt"
	"os"
	"sync"
)

var once sync.Once
var _config *config

type config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func GetConfig() *config {
	once.Do(func() {
		_config = &config{
			Host:     "db",
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Database: os.Getenv("POSTGRES_DB"),
			SSLMode:  os.Getenv("SSL"),
		}
		fmt.Println(_config)
	})
	return _config
}

func (c *config) GetConnString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.Database, c.SSLMode)
}
