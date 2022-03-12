package pgsql

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

const (
	KDefaultTimeout = 30 * time.Second
)

type Config struct {
	Name        string `json:"name,omitempty"`
	Environment string `json:"environment,omitempty"`
	DSN         string `json:"dsn,omitempty"`
	Active      int    `json:"active,omitempty"`
	Idle        int    `json:"idle,omitempty"`
	Lifetime    int    `json:"lifetime,omitempty"` // Connection's lifetime in seconds
}

func getConfigFromEnv() *Config {
	configKey := "Postgres"
	config := &Config{}
	if err := viper.UnmarshalKey(configKey, &config); err != nil {
		err := fmt.Errorf("not found config name with env %q for Postgres with error: %+v", configKey, err)
		panic(err)
	}

	if config.DSN == "" {
		err := fmt.Errorf("not found dns env %q for Postgres", configKey)
		panic(err)
	}

	if config.Name == "" {
		config.Name = "immaster"
	}

	if config.Environment == "" {
		config.Environment = "dev"
	}

	if config.Active == 0 {
		config.Active = 50
	}

	if config.Idle == 0 {
		config.Idle = 50
	}

	if config.Lifetime == 0 {
		config.Lifetime = 5 * 60
	}
	return config
}
