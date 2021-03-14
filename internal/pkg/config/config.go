package config

import (
	"os"
	"strconv"
)

func GetEnvInt(key string) int {
	val := os.Getenv(key)
	if val == "" {
		panic("Empty config")
	}

	intVal, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}

	return intVal
}
