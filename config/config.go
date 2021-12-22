package config

import (
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type ThirdAppAdapter struct {
	Database interface{}
	Redis    interface{}
}

type ServerConfig struct {
	GROUP_API      string
	PORT           string
	READTIMEOUT    time.Duration
	WRITETIMEOUT   time.Duration
	MAXHEADERBYTES int
}

func ServerConfigs() *ServerConfig {
	return &ServerConfig{
		GROUP_API:      "/api/v1",
		PORT:           ":3000",
		READTIMEOUT:    5,
		WRITETIMEOUT:   7,
		MAXHEADERBYTES: 1 << 20,
	}
}

type DatabaseConfig struct {
	DB_URI string
}

func DatabaseConfigs() *DatabaseConfig {
	return &DatabaseConfig{
		DB_URI: os.Getenv("DB_URI"),
	}
}

type RedisConfig struct {
	REDIS_URI        string
	REDIS_PASS       string
	REDIS_DEFAULT_DB int
}

func RedisConfigs() *RedisConfig {
	return &RedisConfig{
		REDIS_URI:        os.Getenv("REDIS_URI"),
		REDIS_PASS:       os.Getenv("REDIS_PASS"),
		REDIS_DEFAULT_DB: 0,
	}
}

func PrivateKey() (key string) {
	key = os.Getenv("PRIVATE_KEY")
	if key == "" {
		raw, err := ioutil.ReadFile("/etc/machine-id")

		if err != nil {
			return ""
		}

		key = string(raw)

		if lines := strings.Split(key, "\n"); len(lines) == 0 {
			return ""
		} else {
			key = lines[0]
		}
	}

	return
}
