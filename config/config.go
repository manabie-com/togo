package config

import "time"

type (
	Config struct {
		State      string
		RestfulAPI struct {
			Host string `yaml:"host"`
			Port string `yaml:"port"`
		} `yaml:"restful_api"`
		Store string `yaml:"store"`
		DBs   struct {
			SQLite struct {
				DataSourceName string `yaml:"data_source_name"`
			} `yaml:"sqlite"`
			Postgres struct {
				Host     string `yaml:"host"`
				Port     string `yaml:"port"`
				Database string `yaml:"database"`
				Username string `yaml:"username"`
				Password string `yaml:"password"`
			} `yaml:"postgres"`
		} `yaml:"dbs"`
		JWTKey   string        `yaml:"jwt_key"`
		SSExpire time.Duration `yaml:"ss_expire"`
	}
)
