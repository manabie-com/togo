package config

import (
	"fmt"
	"os"
	"sync"
)

var Jwt = os.Getenv("JWT_KEY")

type pathEnum struct {
	PathLogin string
	PathTasks string
}

var PathConfig = &pathEnum{
	PathLogin: "/login",
	PathTasks: "/tasks",
}

type dbType struct {
	Postgres string
	Sqlite   string
}

var DBType = &dbType{
	Postgres: "postgres",
	Sqlite:   "sqlite3",
}

type postgresDBConfig struct {
	DBHost     string
	DBName     string
	DBPort     string
	DBUsername string
	DBPassword string
	SSLModel   string
}

var PostgresDBConfig *postgresDBConfig
var lock = &sync.Mutex{}

func GetPostgresDBConfig() *postgresDBConfig {
	if PostgresDBConfig == nil {
		lock.Lock()
		defer lock.Unlock()
		if PostgresDBConfig == nil {
			PostgresDBConfig = &postgresDBConfig{}
			PostgresDBConfig.DBHost = os.Getenv("DB_HOST")
			PostgresDBConfig.DBName = os.Getenv("POSTGRES_DB")
			PostgresDBConfig.DBPort = os.Getenv("POSTGRES_PORT")
			PostgresDBConfig.DBUsername = os.Getenv("POSTGRES_USER")
			PostgresDBConfig.DBPassword = os.Getenv("POSTGRES_PASSWORD")
			PostgresDBConfig.SSLModel = os.Getenv("POSTGRES_SSL")
		}
	}
	return PostgresDBConfig
}

func (p *postgresDBConfig) ToString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", p.DBHost, p.DBPort, p.DBUsername, p.DBPassword, p.DBName, p.SSLModel)
}

type ErrorInfo struct {
	Err        error
	StatusCode int
}

func (e *ErrorInfo) Error() string {
	return e.Err.Error()
}
