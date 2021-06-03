package app

import "os"

var POSTGRESQL_URL string

var (
	ConfigPostgres = map[string]string{
		"host": "127.0.0.1",
		"name": "togo",
		"port": "5434",
		"type": "postgres",
		"user": "postgres",
		"pass": "postgres",
	}
)

const (
	APP_ENV             = "APP_ENV"
	COOKIE_ACCESS_TOKEN = "_acc"
)

func Initialize() {
	if os.Getenv(APP_ENV) == "docker" {
		POSTGRESQL_URL = "postgres://docker:docker@postgres:5432/togo?sslmode=disable"
	} else {
		POSTGRESQL_URL = "postgres://postgres:postgres@localhost:5434/togo?sslmode=disable"
	}

}
