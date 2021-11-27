package config

import (
	"bytes"
	"github.com/spf13/viper"
	"main/internal/logger"
	"strings"
)

var defaultConfig = []byte(`
http_address: 9000
storage_type: InMemory
`)

type Config struct {
	Base    `mapstructure:",squash"`
	Storage `mapstructure:",squash"`
}

type Base struct {
	HTTPAddress int `yaml:"http_address" mapstructure:"http_address"`
}

type Storage struct {
	StorageType string `yaml:"storage_type" mapstructure:"storage_type"`
}

var (
	log = logger.New()
)

func Load() *Config {
	var cfg = &Config{}

	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer(defaultConfig))
	if err != nil {
		log.Fatal("Failed to read viper config", logger.Error(err))
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal("Failed to unmarshal config", logger.Error(err))
	}

	log.Info("Config loaded", logger.Object("config", cfg))
	return cfg
}
