package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

var cf *Configuration

type Configuration struct {
	EnvironmentPrefix string `mapstructure:"environment_prefix"`
	ServerPort        int    `mapstructure:"server_port"`
	DbConnection      string `mapstructure:"db_connection"`
	DefaultPageNum    int    `mapstructure:"default_page_num"`
	DefaultPageLimit  int    `mapstructure:"default_page_limit"`
}

func GetConfig() *Configuration {
	return cf
}

// InitFromFile init Config file
func InitFromFile(path string) *Configuration {
	if path == "" {
		viper.AddConfigPath("config")
		viper.SetConfigType("toml")
		viper.SetConfigName("config")
	} else {
		viper.SetConfigFile(path)
	}
	basePath, _ := os.Getwd()
	viper.AutomaticEnv()
	viper.SetEnvPrefix("")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil && !os.IsNotExist(err) {
		log.Fatalf("Config file not found: %v", err)
	}
	viper.Set("base_path", basePath)
	if err := viper.Unmarshal(&cf); err != nil {
		log.Fatalf("covert to struct: %v", err)
	}
	if path == "" {
		fmt.Printf("File config used %s\n \n", viper.ConfigFileUsed())
		dataPrinf, _ := json.Marshal(cf)
		fmt.Printf("Config: %s\n \n", string(dataPrinf))
	}
	return cf
}
