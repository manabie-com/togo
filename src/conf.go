package server

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Environment string `json:"Env,omitempty"`
	Host        string `json:"Host,omitempty"`
	Port        string `json:"Port,omitempty"`
}

func getConfigFromEnv() *Config {
	confKey := "Server"
	conf := &Config{}
	if err := viper.UnmarshalKey(confKey, &conf); err != nil {
		err = fmt.Errorf("not found config name with env %q for Redis with error: %+v", confKey, err)
		panic(err)
	}
	return conf
}
