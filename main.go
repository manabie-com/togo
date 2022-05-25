package main

import (
	"github.com/japananh/togo/common"
	"github.com/japananh/togo/component/tokenprovider"
	"github.com/japananh/togo/server"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// //go:embed migrations/*.sql
//var embedMigrations embed.FS

func main() {
	// Load config
	config := common.NewConfig()
	if err := config.Load("."); err != nil {
		log.Fatalln("cannot load config from env file", err)
	}

	dbConn, err := gorm.Open(mysql.Open(config.DBConnectionURL()), &gorm.Config{})
	if err != nil {
		log.Fatalln("cannot open database connection", err)
	}

	// create token configs
	tokenConfig, err := tokenprovider.NewTokenConfig(config.AtExpiry(), config.RtExpiry())
	if err != nil {
		log.Fatalln("cannot create token config", err)
	}

	s := server.Server{
		Port:        config.AppPort(),
		AppEnv:      config.AppEnv(),
		SecretKey:   config.SecretKey(),
		DBConn:      dbConn,
		TokenConfig: tokenConfig,
		ServerReady: make(chan bool),
	}

	go func() {
		<-s.ServerReady
		close(s.ServerReady)
	}()

	s.Start()
}

//func runDBMigrations(db *gorm.DB, testDb *gorm.DB) error {
//	sqlDB, err := db.DB()
//	if err != nil {
//		return err
//	}
//
//	sqlTestDB, err := testDb.DB()
//	if err != nil {
//		return err
//	}
//
//	goose.SetBaseFS(embedMigrations)
//
//	if err := goose.SetDialect("mysql"); err != nil {
//		return err
//	}
//
//	if err := goose.Up(sqlDB, "migrations"); err != nil {
//		return err
//	}
//
//	if err := goose.Up(sqlTestDB, "migrations"); err != nil {
//		return err
//	}
//
//	return nil
//}
