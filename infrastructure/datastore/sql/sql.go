package sql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	DriverName string
	Host       string
	Port       string
	Username   string
	Password   string
	DBName     string
	SSLMode    string
}

func (c Config) String() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.Username, c.Password, c.DBName, c.SSLMode,
	)
}

type ConfigOption func(*Config)

func WithDriverName(driverName string) ConfigOption {
	return func(c *Config) {
		c.DriverName = driverName
	}
}

func WithHost(host string) ConfigOption {
	return func(c *Config) {
		c.Host = host
	}
}

func WithPort(port string) ConfigOption {
	return func(c *Config) {
		c.Port = port
	}
}

func WithUsername(username string) ConfigOption {
	return func(c *Config) {
		c.Username = username
	}
}

func WithPassword(password string) ConfigOption {
	return func(c *Config) {
		c.Password = password
	}
}

func WithDBName(dbname string) ConfigOption {
	return func(c *Config) {
		c.DBName = dbname
	}
}

func Connect(options ...ConfigOption) (*sqlx.DB, error) {
	c := &Config{
		Host:    "localhost",
		Port:    "5432",
		DBName:  "test_db",
		SSLMode: "disable",
	}

	for _, o := range options {
		o(c)
	}

	dsn := c.String()

	db, err := sqlx.Open(c.DriverName, dsn)
	if err != nil {
		return nil, fmt.Errorf("connecting [%s]: %w", dsn, err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
