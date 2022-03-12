package log

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Path     string `json:"Path,omitempty"`
	FileName string `json:"FileName,omitempty"`
}

var LogConfigData *Config

func getConfigFromEnv() *Config {
	confKey := "Log"
	conf := &Config{}
	if err := viper.UnmarshalKey(confKey, conf); err != nil {
		err = fmt.Errorf("not found config name with env %q for Redis with error: %+v", confKey, err)
		panic(err)
	}
	return conf
}
