package config

import (
	"os"
	"strconv"
)

type ToGo struct {
	ServicePort int

	DBDriver string
	DSN      string
}

func Load() *ToGo {
	return &ToGo{
		ServicePort: envInt("TOGO_SERVICE_PORT", 9000),

		DBDriver: env("TOGO_DB_DRIVER", ""),
		DSN:      env("TOGO_DSN", ""),
	}
}

func env(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}

func envInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	v, err := strconv.Atoi(value)
	if err == nil {
		return v
	}

	return defaultValue
}
