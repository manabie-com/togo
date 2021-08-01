package config

import "database/sql"

type ServerConfig struct {
	*Config
	*sql.DB
}
