package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite" // Sqlite driver based on GGO
	"gorm.io/gorm"
	"os"
	"os/signal"
	"togo-thdung002/config"
	"togo-thdung002/controllers"
)

func main() {
	localconf, _ := config.LoadConfig("./config/config.json")
	db, err := loadDatabase(localconf)
	if err != nil {
		panic(err)
	}
	svc := controllers.NewController(localconf, db)
	svc.Start()
	defer svc.Stop()
	svc.Load()
	go svc.ListenAndServe()
	log.Println("\nServer started.")
	sigHandler := func() {
		signChan := make(chan os.Signal, 1)
		signal.Notify(signChan, os.Interrupt)
		sig := <-signChan
		log.Println("Cleanup processes started by ", sig, " signal")
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		svc.Shutdown(ctx)
		os.Exit(1)
	}

	sigHandler()

}

func loadDatabase(cfg *config.Config) (*gorm.DB, error) {
	var err error
	db, err := gorm.Open(sqlite.Open(cfg.API.DBAddress), &gorm.Config{})
	if err == nil {
		fmt.Println("connect database successful", cfg.API.DBAddress)
		return db, nil
	}
	return nil, err
}
