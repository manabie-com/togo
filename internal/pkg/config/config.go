package config

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type (
	Server struct {
		Addr              string         `mapstructure:"address"`
		Port              string         `mapstructure:"port"`
		ReadTimeout       *time.Duration `mapstructure:"read_timeout"`
		WriteTimeout      *time.Duration `mapstructure:"write_timeout"`
		ReadHeaderTimeout *time.Duration `mapstructure:"read_header_timeout"`
		ShutdownTimeOut   *time.Duration `mapstructure:"shutdown_timeout"`
	}

	DB struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
		Driver   string `mapstructure:"driver"`
	}

	Log struct {
		Level   string `mapstructure:"level"`
		Encoder string `mapstructure:"encode"`
		File    string `mapstructure:"file"`
	}
	Config struct {
		Server Server `mapstructure:"server"`
		DB     DB     `mapstructure:"db"`
		Log    Log    `mapstructure:"log"`
		Env    string `mapstructure:"env"`
	}
)

const (
	envPrefix      = "TOGO"
	configPathFlag = "configs"
	fileType       = "yaml"
	baseFilename   = "config-base"
	configFileName = "config"
)

var (
	All = &Config{}
)

func Load() {
	pflag.StringP(configPathFlag, "c", "./configs", "Load configuration variables from dir")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(fmt.Sprintf("bind flag %v", err))

	}

	// read config variables from env with prefix
	// ex: TOGO_DB_URL
	viper.AutomaticEnv()
	viper.AllowEmptyEnv(true)
	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// read config variables from file
	viper.AddConfigPath(viper.GetString(configPathFlag))
	viper.SetConfigType(fileType)

	// set base config
	viper.SetConfigName(baseFilename)
	if err := viper.ReadInConfig(); err != nil {
		if errors.As(err, &viper.ConfigFileNotFoundError{}) {
			log.Println("not found")
		}
		log.Panicf("read config-base %v", err)
	}

	// set config for each env (dev, staging, prod)
	viper.SetConfigName(configFileName)

	if err := viper.MergeInConfig(); err != nil {
		log.Panicf("merge config %v", err)
	}

	if err := viper.Unmarshal(&All); err != nil {
		log.Panicf("unmarshal config %v", err)
	}
}
