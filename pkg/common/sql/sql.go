package sql

import (
	"fmt"
	"log"
)

type ConfigPostgres struct {
	Protocol string `yaml:"protocol"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	SSLMode  string `yaml:"sslmode"`
	Timeout  int    `yaml:"timeout"`
}

func (c *ConfigPostgres) ConnectionString() (driver string, connStr string) {
	sslmode := c.SSLMode
	if c.SSLMode == "" {
		sslmode = "disable"
	}
	if c.Timeout == 0 {
		c.Timeout = 15
	}

	switch c.Protocol {
	case "":
	case "postgres":
		connStr = fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v connect_timeout=%v", c.Host, c.Port, c.Username, c.Password, c.Database, sslmode, c.Timeout)
	default:
		log.Fatalf("postgres: Invalid protocol %s", c.Protocol)
	}
	return c.Driver(), connStr
}

func (c *ConfigPostgres) Driver() (driver string) {

	switch c.Protocol {
	case "":
	case "postgres":
		driver = "postgres"
	default:
		log.Fatalf("unsupported protocol %s", c.Protocol)
	}
	return
}