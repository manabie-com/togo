package mysql

import (
	"context"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/khangjig/togo/client/logger"
	"github.com/khangjig/togo/config"
)

var db *gorm.DB

func init() {
	var (
		err error
		cfg = config.GetConfig()
	)

	connectionString := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySQL.User,
		cfg.MySQL.Pass,
		cfg.MySQL.Host,
		cfg.MySQL.Port,
		cfg.MySQL.DBName,
	)

	db, err = gorm.Open(mysql.New(mysql.Config{DSN: connectionString}), &gorm.Config{})
	if err != nil {
		logger.GetLogger().Fatal(err.Error())
	}

	db = db.Debug()

	logger.GetLogger().Info("Connected mysql db")
}

func GetClient(ctx context.Context) *gorm.DB {
	return db.Session(&gorm.Session{})
}
