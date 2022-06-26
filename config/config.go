package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var Config *Configuration

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
	Cors     CorsConfiguration
}

type DatabaseConfiguration struct {
	Driver       string
	Dbname       string
	Username     string
	Password     string
	Host         string
	Port         string
	Charset      string
	MaxLifetime  int `mapstructure:"max_lifetime"`
	MaxOpenConns int `mapstructure:"max_open_conns"`
	MaxIdleConns int `mapstructure:"max_idle_conns"`
}

type ServerConfiguration struct {
	ServerPort string `mapstructure:"server_port"`
	GinMode    string `mapstructure:"gin_mode"`
}

type CorsConfiguration struct {
	AllowedOrigins []string `mapstructure:"allowed_origins"`
	AllowedMethods []string `mapstructure:"allowed_methods"`
	AllowedHeaders []string `mapstructure:"allowed_headers"`
}

type CacheConfiguration struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// SetupDB initialize configuration
func Setup(configPath string) error {
	var configuration *Configuration

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config dir: %s, %v", configPath, err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		return fmt.Errorf("unable to decode into struct, %v", err)
	}

	Config = configuration

	return nil
}

// GetConfig helps you to get configuration data
func GetConfig() *Configuration {
	return Config
}
