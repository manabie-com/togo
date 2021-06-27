package clients

import (
	"github.com/avast/retry-go"
	"gorm.io/gorm"
	"time"
	"gorm.io/driver/postgres"
)

type PSQLConfig struct {
	DSN                   string `envconfig:"POSTGRES_DSN" required:"true"`
	ConnMaxLifeTimeSecond int64  `envconfig:"POSTGRES_CONN_MAX_LIFE_TIME_SECOND" default:"300"`
}

func InitPSQLDB(cfg PSQLConfig) (*gorm.DB, error) {
	const (
		maxAttempts   = 10
		delaySeconds  = 6
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
	return gormDB, nil
}