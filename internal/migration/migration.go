package migration

import (
	"fmt"

	"github.com/TrinhTrungDung/togo/config"
	"github.com/TrinhTrungDung/togo/pkg/db"
	"github.com/TrinhTrungDung/togo/pkg/migration"
	"gopkg.in/gormigrate.v1"
)

func Run() (resErr error) {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	db, err := db.New(fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=%s", cfg.DbDialect, cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName, cfg.DbSslMode), false)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				resErr = fmt.Errorf("%s", x)
			case error:
				resErr = x
			default:
				resErr = fmt.Errorf("Unknown error: %+v", x)
			}
		}
	}()

	// Create migrations table to store migration version history
	initSQL := "CREATE TABLE IF NOT EXISTS migrations (id VARCHAR(255) PRIMARY KEY);"
	if err := db.Exec(initSQL).Error; err != nil {
		return err
	}

	migration.Run(db, []*gormigrate.Migration{
		{},
	})

	return nil
}
