package config

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

type Config struct {
	API struct {
		Listen    string `json:"listen"`
		DBAddress string `json:"db"`
	} `json:"api"`
}

func getConfigFromFile(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func decodeLocalConf(data []byte) (*Config, error) {
	config := &Config{}

	// try to decode json first and yaml in the next step
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	// validate config
	//config.Prefix = strings.TrimSuffix(config.Prefix, "/")
	return config, nil
}

func LoadConfig(filename string) (*Config, error) {

	log.Info("Loading configuration from ", filename)

	confData, jerr := getConfigFromFile(filename)
	if jerr != nil {
		log.Fatal("Failed to load configuration: ", jerr)
		return nil, jerr
	}

	return decodeLocalConf(confData)
}
