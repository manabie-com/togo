package util

import (
	"github.com/manabie-com/togo/internal/config"
	"regexp"
)

const DriverPostgres = "postgres"
const DriverSQLite3 = "sqlite3"

func GetConnectionString(conf config.DB) string {
	switch conf.DriverName {
	case DriverPostgres:
		return conf.Postgres.ToConnectionString()
	case DriverSQLite3:
		return conf.SQLite3.DataSourceName
	}
	return ""
}

var sqlRegex = regexp.MustCompile(`\$[1-9][0-9]*`)

func PrepareQuery(driverName, query string) string {
	switch driverName {
	case DriverPostgres:
		return query
	case DriverSQLite3:
		return sqlRegex.ReplaceAllString(query, "?")
	}
	return query
}
