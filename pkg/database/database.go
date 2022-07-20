package database

import (
	"database/sql"
	"fmt"
	"time"
)

// DBConfig used to set config for database.
type DBConfig interface {
	String() string
	DSN() string
}

type MySQLConfig struct {
	Host     string `json:"host" mapstructure:"host" yaml:"host"`
	Database string `json:"database" mapstructure:"database" yaml:"database"`
	Port     int    `json:"port" mapstructure:"port" yaml:"port"`
	Username string `json:"username" mapstructure:"username" yaml:"username"`
	Password string `json:"password" mapstructure:"password" yaml:"password"`
	Options  string `json:"options" mapstructure:"options" yaml:"options"`
}

// DSN returns the Domain Source Name.
func (c MySQLConfig) DSN() string {
	options := c.Options
	if options != "" {
		if options[0] != '?' {
			options = "?" + options
		}
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s%s",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
		options)
}

// String returns MySQL connection URI.
func (c MySQLConfig) String() string {
	return fmt.Sprintf("mysql://%s", c.DSN())
}

func NewMySQLDatabase(config DBConfig) *sql.DB {
	dbPool, err := sql.Open("mysql", config.DSN())
	if err != nil {
		panic(err)
	}

	dbPool.SetMaxIdleConns(5)
	dbPool.SetMaxOpenConns(25)
	dbPool.SetConnMaxLifetime(time.Hour)

	return dbPool
}
