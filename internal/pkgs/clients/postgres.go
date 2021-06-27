package clients

import (
	"github.com/avast/retry-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type PSQLConfig struct {
	DSN                   string `envconfig:"POSTGRES_DSN" required:"true" default:"host=localhost user=togo password=ad34a$dg dbname=manabie_togo port=5432 sslmode=disable timezone=UTC"`
	ConnMaxLifeTimeSecond int64  `default:"300"`
	ConnMaxIdleTime       int    `default:"10"`
	MaxOpenConns          int    `default:"100"`
}

func InitPSQLDB(cfg PSQLConfig) (*gorm.DB, error) {
	const (
		maxAttempts  = 10
		delaySeconds = 6
	)
	var gormDB *gorm.DB
	err := retry.Do(
		func() error {
			db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})
			gormDB = db
			return err
		},
		retry.DelayType(retry.FixedDelay),
		retry.Attempts(maxAttempts),
		retry.Delay(delaySeconds*time.Second),
		retry.LastErrorOnly(true),
	)
	if err != nil {
		return nil, err
	}
	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, err
	}
	if cfg.ConnMaxLifeTimeSecond > 0 {
		sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifeTimeSecond) * time.Second)
	}
	if cfg.ConnMaxIdleTime > 0 {
		sqlDB.SetConnMaxIdleTime(time.Duration(cfg.ConnMaxLifeTimeSecond) * time.Second)
	}
	if cfg.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	}

	return gormDB, nil
}
