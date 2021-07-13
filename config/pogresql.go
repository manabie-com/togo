package config

import (
	"fmt"
	"time"
)

// PostgreSQLConfig schema
type PostgreSQLConfig struct {
	Host            string
	Database        string
	Port            int
	Username        string
	Password        string
	Options         string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

// String return PostgreSQL connection url
func (m PostgreSQLConfig) String() string {
	return fmt.Sprintf("postgres://%s", m.DSN())
}

// DSN return Data Source Name
func (m PostgreSQLConfig) DSN() string {
	return fmt.Sprintf("%s:%s@%s:%d/%s%s", m.Username, m.Password, m.Host, m.Port, m.Database, m.Options)
}
