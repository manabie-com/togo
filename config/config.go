package config

import (
	"time"

	"github.com/spf13/viper"
)

func init() {
	load()
}

var (
	JWT             JWTConfig
	HTTPPort        int
	PostgreSQL      PostgreSQLConfig
	MigrationFolder string
)

type (
	JWTConfig struct {
		Key       string
		ExpiresIn time.Duration
	}
)

func load() {
	v := viper.New()
	v.SetConfigType("json")
	v.AutomaticEnv()

	v.SetDefault("HTTP_PORT", 5050)
	v.SetDefault("JWT_KEY", "wqGyEBBfPK9w3Lxw")
	v.SetDefault("JWT_EXPIRES_IN", 15*time.Minute) // 15 minutes

	HTTPPort = v.GetInt("HTTP_PORT")
	JWT = JWTConfig{
		Key:       v.GetString("JWT_KEY"),
		ExpiresIn: v.GetDuration("JWT_EXPIRES_IN"),
	}

	PostgreSQL = PostgreSQLConfig{
		Host:     v.GetString("POSTGRE_HOST"),
		Username: v.GetString("POSTGRES_USERNAME"),
		Password: v.GetString("POSTGRES_PASSWORD"),
		Database: v.GetString("POSTGRE_DATABASE"),
		Port:     v.GetInt("POSTGRES_PORT"),
		Options:  "?sslmode=disable",

		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxLifetime: 8 * time.Minute,
	}

	maxIdleConns := v.GetInt("POSTGRES_MAX_IDLE_CONNS")
	maxOpenConns := v.GetInt("POSTGRES_MAX_OPEN_CONNS")
	connMaxLifetime := v.GetInt("POSTGRES_CONN_MAX_LIFETIME")

	if maxIdleConns > 0 {
		PostgreSQL.MaxIdleConns = maxIdleConns
	}
	if maxOpenConns > 0 {
		PostgreSQL.MaxOpenConns = maxOpenConns
	}
	if connMaxLifetime > 0 {
		PostgreSQL.ConnMaxLifetime = time.Duration(connMaxLifetime)
	}
	MigrationFolder = "file://migrations"
}
