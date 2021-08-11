package cmd

import "github.com/manabie-com/togo/common/database"

type HostConfig struct {
	Host string
	Port int
}

type TokenConfig struct {
	Key     string
	Timeout int
}

type AppConfig struct {
	Host  *HostConfig
	Db    *database.Config
	Token *TokenConfig
}
