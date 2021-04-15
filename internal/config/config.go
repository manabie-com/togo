package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitializeConfig(configPath string) (Config, error) {
	vp := viper.New()
	vp.AddConfigPath(configPath)
	vp.SetConfigType("env")
	vp.SetConfigFile(".env")
	err := vp.ReadInConfig()
	if err != nil {
		return Config{}, err
	}
	var c Config
	err = vp.Unmarshal(&c)
	if err != nil {
		return Config{}, err
	}
	return c, nil
}

type Config struct {
	JWTKey string `mapstructure:"TOGO_JWT_KEY"`
	DB     DB     `mapstructure:",squash"`
}

type DB struct {
	DriverName string   `mapstructure:"TOGO_DB_DRIVER_NAME"`
	SQLite3    SQLite3  `mapstructure:",squash"`
	Postgres   Postgres `mapstructure:",squash"`
}

type SQLite3 struct {
	DataSourceName string `mapstructure:"TOGO_DB_SQLITE3_DATA_SOURCE_NAME"`
}

type Postgres struct {
	Host     string `mapstructure:"TOGO_DB_POSTGRES_HOST"`
	Port     string `mapstructure:"TOGO_DB_POSTGRES_PORT"`
	Name     string `mapstructure:"TOGO_DB_POSTGRES_NAME"`
	UserName string `mapstructure:"TOGO_DB_POSTGRES_USERNAME"`
	Password string `mapstructure:"TOGO_DB_POSTGRES_PASSWORD"`
	SSLMode  string `mapstructure:"TOGO_DB_POSTGRES_SSL_MODE"`
}

func (p Postgres) ToConnectionString() string {
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", p.Host, p.Port, p.Name, p.UserName, p.Password, p.SSLMode)
}
