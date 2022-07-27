package config

import (
	"bytes"
	"encoding/json"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	DB DBConfig `mapstructure:"DB" json:"db"`
}

type DBConfig struct {
	Driver          string `mapstructure:"DRIVER" json:"driver"`
	Source          string `mapstructure:"SOURCE" json:"source"`
	MigrationFolder string `mapstructure:"MIGRATION_FOLDER" json:"migration_folder"`
}

func loadDefaultConfig() *Config {
	return &Config{
		DB: loadDefaultDBConfig(),
	}
}

func loadDefaultDBConfig() DBConfig {
	return DBConfig{
		Driver:          "postgres",
		Source:          "postgresql://admin:pass@localhost:5432/mnb?sslmode=disable",
		MigrationFolder: "file://migrations",
	}
}

func Load() (*Config, error) {
	c := loadDefaultConfig()

	viper.SetConfigType("json")
	configBuffer, err := json.Marshal(c)

	if err != nil {
		return nil, err
	}

	err = viper.ReadConfig(bytes.NewBuffer(configBuffer))
	if err != nil {
		return nil, err
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()
	err = viper.Unmarshal(c)
	return c, err
}
