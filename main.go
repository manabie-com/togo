package main

import (
	"context"
	"flag"
	"fmt"
	"runtime"
	"sync"

	"mini_project/server"

	"mini_project/config"
	"mini_project/db"
)

var (
	configFile = flag.String("config", "", "Config file")
	purgeDB    = flag.Bool("purge-db", false, "Purge Your DB")
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	flag.Parse()

	// load config
	dbconfigPath := "config/db_config.yml"
	fmt.Println("configPath", dbconfigPath)
	if *configFile != "" {
		dbconfigPath = *configFile
	}
	err := config.NewConfig(dbconfigPath)
	if err != nil {
		panic(err)
	}

	dburl := config.GetDbUrl()
	// fmt.Println("db url : ", dburl)
	if *purgeDB {
		db.PurgeDB(dburl)
	}

	if err := db.CreateDatabase(dburl); err != nil {
		fmt.Println("could not found database")
		panic(err)
	}

	// load public key
	config.InitJWT()

	// fmt.Println("db url : ", dburl)

	ctx := context.Background()

	wg.Add(1)
	go server.StartGRPC(&wg, ctx, dburl)
	wg.Add(1)
	go server.StartHTTP(&wg) // support restapi

	wg.Wait()
}
