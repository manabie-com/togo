package services

import "fmt"

const (
	limitCacheKeyPrefix         = "limit"
)
const (
	redisHost = "redis:6379"
)
const (
	postgresHost = "postgres"
	port         = 5432
	user         = "dev"
	password     = "passwd"
	dbname       = "togo"
)

var (
	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		postgresHost, port, user, password, dbname)
)
