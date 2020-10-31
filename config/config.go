package config

import (
	"io/ioutil"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

// Config structure
type (
	Config struct {
		AppEnv    string     `yaml:"appEnv" envconfig:"APP_ENV"`
		Service   string     `yaml:"service" envconfig:"SERVICE"`
		TodoStore TodoConfig `yaml:"todoStore"`
		JWTKey    string     `yaml:"JWTKey" envconfig:"JWT_KEY"`
		LogLevel  string     `yaml:"logLevel" envconfig:"LOG_LEVEL"`
		Address   string     `yaml:"address" envconfig:"ADDRESS"`
	}
	// TodoConfig ...
	TodoConfig struct {
		LDB TodoLDBConfig `yaml:"ldb"` // sqlite
	}
	// TodoLDBConfig ...
	TodoLDBConfig struct {
		Path string `yaml:"ldb_path" envconfig:"LDB_PATH"`
	}
)

//LoadConfigFile load default config from file
func LoadConfigFile(path string) Config {
	c := Config{}
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return c
	}
	if err = yaml.Unmarshal(content, &c); err != nil {
		return c
	}
	return c
}

// LoadConfigEnv load config from environment variables
func LoadConfigEnv(c *Config) {
	if err := envconfig.Process("", c); err != nil {
		panic(err)
	}
}
