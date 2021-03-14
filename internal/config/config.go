package config

import "time"

type RedisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

//FIXME: support for postgreDB
type DatabaseConfig struct {
	DriverName     string `json:"driver_name"`
	DataSourceName string `json:"data_source_name"`
}
type Config struct {
	JwtKey   string         `json:"jwt_key"`
	Redis    RedisConfig    `json:"redis"`
	Database DatabaseConfig `json:"database"`
	TokenTIL time.Duration            `json:"token_til"` // Minutes
}
