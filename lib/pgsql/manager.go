package pgsql

import (
	"context"
	"fmt"

	"github.com/HoangMV/todo/lib/log"
	"github.com/jmoiron/sqlx"
)

var (
	client *SQLClient
)

// default value env key is "Postgres";
// if configKeys was set, key env will be first value (not empty) of this;
func Install() {
	config := getConfigFromEnv()

	client = NewSqlxDB(config)
	if client == nil {
		err := fmt.Errorf("InstallPostgresSQLClientManager - NewSqlxDB {%v} error", config)
		log.Get().Errorf("InstallPostgresSQLClientManager - Error: %v", err)

		panic(err)
	}

	if config.Name == "" {
		err := fmt.Errorf("InstallPostgresSQLClientManager - config error: config.Name is empty")
		log.Get().Errorf("InstallPostgresSQLClientManager - Error: %v", err)

		panic(err)
	}
	fmt.Println("====> PostgreSql Install Done!")
}

func Get() *sqlx.DB {
	return client.DB
}

func GetDefaultContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), KDefaultTimeout)
}
