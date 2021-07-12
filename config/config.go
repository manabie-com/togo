package config

import (
	"fmt"
	"github.com/manabie-com/togo/internal/pkg/logger"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
	"time"
)

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

func Load(state *string) (*Config, error) {
	cfgPath := fmt.Sprintf("%v/config/config.%v.yml", RootDir(), *state)
	f, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		logger.Errorf("Fail to open configurations file: %v", err)
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(f, &cfg)
	if err != nil {
		logger.Errorf("Fail to decode configurations file: %v", err)
		return nil, err
	}
	cfg.State = *state
	return &cfg, nil
}

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
