package config

import (
	"database/sql"
	"github.com/go-redis/redis/v8"
)

type ServerConfig struct {
	*Config
	*sql.DB
	Redis *redis.Client
}
