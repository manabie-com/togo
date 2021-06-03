package utils

import (
	"os"
	"strconv"
)

// Retrieve env key, return default value if value of env is empty
func GetStringEnv(envKey string, defaultVal string) string {
	if val := os.Getenv(envKey); val != "" {
		return val
	}
	return defaultVal
}

func GetIntEnv(envKey string, defaultVal int) int {
	if val, err := strconv.Atoi(os.Getenv(envKey)); err == nil {
		return val
	}
	return defaultVal
}
