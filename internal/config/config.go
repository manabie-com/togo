package config

import (
	"bytes"
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"strings"
	"time"
)

type RedisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

//FIXME: support for postgreDB
type DatabaseConfig struct {
	Addr     string `json:"addr"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}
type Config struct {
	JwtKey   string         `json:"jwt_key"`
	Redis    RedisConfig    `json:"redis"`
	Database DatabaseConfig `json:"database"`
	TokenTimeDuration uint `json:"token_time_duration"`  // Minutes
}

func(c*Config) TokenTIL() time.Duration{
	return time.Duration(c.TokenTimeDuration) * time.Minute
}

// Load config from env
func LoadConfigFromEnv(cfg interface{}) error {
	v := viper.New()
	v.SetConfigType("json")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	v.AutomaticEnv()

	cfgJsonBytes, err := json.Marshal(cfg)
	err = v.ReadConfig(bytes.NewReader(cfgJsonBytes))

	err = v.Unmarshal(cfg, func(c *mapstructure.DecoderConfig) {
		c.TagName = "json"
	})

	if err != nil {
		return err
	}

	return nil
}
