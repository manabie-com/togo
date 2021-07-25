package config

import "fmt"

const (
	LimitCacheKeyPrefix = "limit"
)
const (
	RedisHost = "redis:6379"
)
const (
	PostgresHost     = "postgres"
	PostgresPort     = 5432
	PostgreUser      = "dev"
	PostgresPassword = "passwd"
	PostgresDBName   = "togo"
)

var (
	PsqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		PostgresHost, PostgresPort, PostgreUser, PostgresPassword, PostgresDBName)
)
