package env

import (
	"fmt"
	"os"
	"strconv"
)

func getEnv(env string) string {
	if value := os.Getenv(env); value == "" {
		panic(fmt.Sprintf("ENV %s is empty", env))
	} else {
		return value
	}
}

func getIntEnv(env string) int {
	if v, err := strconv.Atoi(getEnv(env)); err == nil {
		return v
	}

	return 0
}

func GetDBHost() string {
	return getEnv("DB_HOST")
}

func GetDBPort() string {
	return getEnv("DB_PORT")
}

func GetDBSSLMode() string {
	return getEnv("DB_SSL_MODE")
}

func GetDBUserName() string {
	return getEnv("DB_USERNAME")
}

func GetDBPassword() string {
	return getEnv("DB_PASSWORD")
}

func GetDBName() string {
	return getEnv("DB_NAME")
}

func GetDBTest() string {
	return getEnv("DB_NAME_TEST")
}

func MaxIdleConns() int {
	return getIntEnv("DB_MAX_IDLE_CONNS")
}

func MaxOpenConns() int {
	return getIntEnv("DB_MAX_OPEN_CONNS")
}

func SecretKeyJWT() string {
	return getEnv("SECRET_KEY_JWT")
}

func ServerPort() int {
	return getIntEnv("SERVER_PORT")
}
