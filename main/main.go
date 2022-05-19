package main

import (
	"flag"
	"fmt"
	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"manabie.com/togo/api"
	"manabie.com/togo/service/tasklimiter"
	"manabie.com/togo/service/user"
	"manabie.com/togo/store/cache"
	"manabie.com/togo/store/persist"
	"os"
	"time"
)

var cfgFile string
var port int

func main() {
	flag.StringVar(&cfgFile, "c", "", "path to your config file")
	flag.IntVar(&port, "port", 8888, "service port")
	flag.Parse()

	if cfgFile == "" {
		log.Fatal("you must define config file")
	}
	initConfigFile(cfgFile)
	if viper.GetInt("server.port") != 0 {
		port = viper.GetInt("server.port")
	}

	cacheImpl, err := cache.New(cache.Config{
		HostAddr:     viper.GetString("redis.hostname"),
		Pass:         viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.db"),
		DialTimeout:  time.Duration(viper.GetInt("redis.dial-timeout")) * time.Second,
		MaxRetries:   viper.GetInt("redis.max-retries"),
		ReadTimeout:  time.Duration(viper.GetInt("redis.read-timeout")) * time.Second,
		WriteTimeout: time.Duration(viper.GetInt("redis.write-timeout")) * time.Second,
	})

	if err != nil {
		log.Fatalf("fail to init cache. Reason: %s", err)
	}

	log.Info("init cache done")
	storeImpl, err := persist.NewStore(persist.Config{
		HostAddr:       viper.GetString("mysql.hostname"),
		DB:             viper.GetString("mysql.dbname"),
		User:           viper.GetString("mysql.username"),
		Pass:           viper.GetString("mysql.password"),
		MaxOpenCnn:     viper.GetInt("mysql.max-open-connection"),
		MaxCnnLifeTime: time.Duration(viper.GetInt("mysql.max-connection-lifetime")) * time.Minute,
	})
	if err != nil {
		log.Fatalf("fail to init storage. Reason: %s", err)
	}

	log.Info("init store done")

	userCrudService := user.New(user.Config{Store: storeImpl})
	taskLimiterService := tasklimiter.New(tasklimiter.Config{
		Store: storeImpl,
		Cache: cacheImpl,
	})

	apiCfg := api.Cfg{
		Port:               port,
		UserCrudService:    userCrudService,
		TaskLimiterService: taskLimiterService,
	}

	api.Init(apiCfg)
	log.Println(api.Run())
}

func initConfigFile(cfgFile string) {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".manabie")
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error when read config file: %s", err)
	} else {
		log.Infof("Using config file: %s", viper.ConfigFileUsed())
	}
}
