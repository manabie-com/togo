package main

import (
	"flag"
	"fmt"
	"github.com/manabie-com/togo/config"
	"github.com/manabie-com/togo/internal/api"
	"github.com/manabie-com/togo/internal/pkg/logger"
	"gopkg.in/yaml.v2"
	"io/ioutil"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var mbLogger logger.Logger

func main() {
	state := flag.String("state", "local", "state of service")
	mbLogger = logger.WithPrefix("main")
	cfg := getConfig(*state)
	initRestfulAPI(cfg)
}

func getConfig(state string) *config.Config {
	cfgPath := fmt.Sprintf("config/config.%v.yml", state)
	f, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		mbLogger.Panicf("Fail to open configurations file: %v", err)
	}

	var cfg config.Config
	err = yaml.Unmarshal(f, &cfg)
	if err != nil {
		mbLogger.Panicf("Fail to decode configurations file: %v", err)
	}
	cfg.State = state
	return &cfg
}

func initRestfulAPI(cfg *config.Config) {
	mbLogger.Info("Start server")
	err := api.CreateAPIEngine(cfg)
	if err != nil {
		mbLogger.Panicf("Fail to listen and server: %v", err)
		return
	}
}
