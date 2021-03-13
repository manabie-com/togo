package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/valonekowd/togo/infrastructure/config"
	"github.com/valonekowd/togo/infrastructure/config/viper"
	"github.com/valonekowd/togo/util/helper"
)

var AppEnvs = []string{"production", "staging", "testing", "development"}

type Config struct {
	AppEnv string
	Server struct {
		Host string
		Port string
	}
	Datastore struct {
		Primary struct {
			DriverName string
			Host       string
			Port       string
			Username   string
			Password   string
			DBName     string
		}
		Cache struct {
			Host string
			Port string
		}
	}
	Auth struct {
		JWT struct {
			Secret    string
			Issuer    string
			Algorithm string
		}
	}
	Validation struct {
		Playground struct {
			TagName string
		}
	}
}

func (c *Config) IsProd() bool {
	return c.AppEnv == "production"
}

func (c *Config) HTTPAddr() string {
	return fmt.Sprintf("%v:%v", c.Server.Host, c.Server.Port)
}

func (c *Config) CacheDSAddr() string {
	return fmt.Sprintf("%v:%v", c.Datastore.Cache.Host, c.Datastore.Cache.Port)
}

func Create() (*Config, error) {
	configPaths := []string{
		filepath.Join(".", "config"),
	}

	c := &Config{}

	var configFiles []*config.File
	{
		f, err := config.NewFile("default", "yml", configPaths)
		if err != nil {
			return nil, err
		}
		configFiles = append(configFiles, f)
	}
	{
		appEnv := os.Getenv("APP_ENV")
		if !helper.StringInSlice(appEnv, AppEnvs) {
			appEnv = "development"
		}

		c.AppEnv = appEnv

		f, err := config.NewFile(c.AppEnv, "yml", configPaths)
		if err != nil {
			return nil, err
		}
		configFiles = append(configFiles, f)
	}

	vc := viper.NewConfiger(
		viper.ConfigerReadFromEnv(true),
		viper.ConfigerFiles(configFiles...),
	)

	return c, vc.Load(c)
}
