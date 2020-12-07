package postgres

import (
	"database/sql"
	"fmt"
	"github.com/HoangVyDuong/togo/pkg/logger"
	"github.com/spf13/viper"
	"time"

	_ "github.com/lib/pq"
)

//NewDatabase new dbPool with config
func NewDatabase() *sql.DB {

	dialect := viper.GetString("database.dialect")
	datasource := viper.GetString("database.url")
	var err error
	dbPool, err := sql.Open(dialect, fmt.Sprintf("%s?sslmode=disable", datasource))
	if err != nil {
		logger.Errorf("Connection to database error: %s", err.Error())
		panic("open database failed")
	}

	dbPool.SetConnMaxLifetime(viper.GetDuration("database.max_lifetime") * time.Second)
	dbPool.SetMaxIdleConns(viper.GetInt("database.max_idle_conns"))
	dbPool.SetMaxOpenConns(viper.GetInt("database.max_open_conns"))

	if err = dbPool.Ping(); err != nil {
		logger.Errorf("Connection to database error: %s", err.Error())
		panic(err.Error())
	}
	return dbPool
}
