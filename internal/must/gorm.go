package must

import (
	"time"

	"github.com/vchitai/l"
	"github.com/vchitai/togo/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	zapgorm "moul.io/zapgorm2"
)

const (
	maxDBIdleConns  = 10
	maxDBOpenConns  = 100
	maxConnLifeTime = 30 * time.Minute
)

var ll = l.New()

func ConnectMySQL(cfg *configs.MySQL) *gorm.DB {
	db, err := gorm.Open(mysql.Open(cfg.FormatDSN()), &gorm.Config{
		Logger: zapgorm.New(ll).LogMode(logger.Silent),
	})
	if err != nil {
		ll.Fatal("Error open mysql", l.Error(err))
	}

	err = db.Raw("SELECT 1").Error
	if err != nil {
		ll.Fatal("Error querying SELECT 1", l.Error(err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		ll.Fatal("Error get sql DB", l.Error(err))
	}
	sqlDB.SetMaxIdleConns(maxDBIdleConns)
	sqlDB.SetMaxOpenConns(maxDBOpenConns)
	sqlDB.SetConnMaxLifetime(maxConnLifeTime)
	return db
}
