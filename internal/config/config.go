package config

import (
	"bytes"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"log"
	"strings"
)

var defaultConfig = []byte(`
environment: T
port: 5050
jwt_key: wqGyEBBfPK9w3Lxw
sqlite: ./data.db
max_todo: 5
postgres:
  host: localhost
  port: 5432
  user:  postgres
  password: 12345
  dbname: togo
  ssl: disable
redis:
  address: localhost:6379
  database_num: 0
  password:
  max_idle: 3
  max_idle_timeout: 300
  read_timeout: 15
  write_timeout: 15
  connect_timeout: 15
`)

type (
	Config struct {
		MaxTodo     int32     `yaml:"max_todo" mapstructure:"max_todo"`
		Environment string    `yaml:"environment" mapstructure:"environment"`
		Port        int       `yaml:"port" mapstructure:"port"`
		JWTKey      string    `yaml:"jwt_key" mapstructure:"jwt_key"`
		SQLite      string    `yaml:"sqlite" mapstructure:"sqlite"`
		Postgres    *Postgres `yaml:"postgres" mapstructure:"postgres"`
		Redis       *Redis    `yaml:"redis" mapstructure:"redis"`
	}

	Postgres struct {
		Host     string `yaml:"host" mapstructure:"host"`
		Port     int    `yaml:"port" mapstructure:"port"`
		User     string `yaml:"user" mapstructure:"user"`
		Password string `yaml:"password" mapstructure:"password"`
		DBName   string `yaml:"dbname" mapstructure:"dbname"`
		SSL      string `yaml:"ssl" mapstructure:"ssl"`
	}

	Redis struct {
		Address                string `yaml:"address" mapstructure:"address"`
		Password               string `yaml:"password" mapstructure:"password"`
		DatabaseNum            int    `yaml:"database_num" mapstructure:"database_num"`
		MaxIdle                int    `yaml:"max_idle" mapstructure:"max_idle"`
		MaxActive              int    `yaml:"max_active" mapstructure:"max_active"`
		MaxIdleTimeout         int    `yaml:"max_idle_timeout" mapstructure:"max_idle_timeout"`
		Wait                   bool   `yaml:"wait" mapstructure:"wait"`
		ReadTimeout            int    `yaml:"read_timeout" mapstructure:"read_timeout"`
		WriteTimeout           int    `yaml:"write_timeout" mapstructure:"write_timeout"`
		ConnectTimeout         int    `yaml:"connect_timeout" mapstructure:"connect_timeout"`
	}
)

func (c *Config) String() string {
	data, _ := yaml.Marshal(c)
	return string(data)
}

func Load() *Config {
	var cfg = &Config{}

	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer(defaultConfig))
	if err != nil {
		log.Fatalf("Failed to read viper config - %s", err)
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("Failed to unmarshal config - %s", err)
	}

	log.Printf("config loaded - %s", cfg.String())
	return cfg
}
