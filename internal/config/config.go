package config

import (
	"encoding/json"
	"os"
)

// Load load a jsonFile into config
func Load(jsonFile string, config interface{}) {
	configFile, _ := os.Open(jsonFile)
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
}
