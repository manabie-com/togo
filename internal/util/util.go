package util

import "os"

func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	} else {
		return defaultVal
	}
}
