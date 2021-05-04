package server

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

// DB is same the sql.DB, but some methods have been provided as asynchronous implementation.
type DataBase struct {
	Postgres *pgxpool.Pool
}

func (db *DataBase) Close() {
	if db.Postgres != nil {
		db.Postgres.Close()
	}
}

func (db *DataBase) InitDatabase() {
	config, err := pgxpool.ParseConfig(Config.DBUrl)
	config.MaxConnIdleTime = time.Duration(Config.DBMaxIdleTime) * time.Second
	config.MaxConnLifetime = time.Duration(Config.DBMaxLifeTime) * time.Second
	config.MinConns = int32(Config.DBMaxConnection)
	config.MaxConns = int32(Config.DBMinConnection)

	db.Postgres, err = pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal(" cant connect to database ")
	}
}
