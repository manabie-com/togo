package infra

import (
	"time"

	"github.com/go-pg/pg"
)

func ProvidePostgres(cfg *AppConfig) (db *pg.DB, cleanup func(), err error) {
	timeoutSecond := cfg.Postgres.ReadTimeout
	if timeoutSecond == 0 {
		timeoutSecond = 60
	}

	db = pg.Connect(&pg.Options{
		Addr:        cfg.Postgres.Address,
		User:        cfg.Postgres.Username,
		Password:    cfg.Postgres.Password,
		Database:    cfg.Postgres.Database,
		ReadTimeout: time.Duration(timeoutSecond) * time.Second,
	})

	if _, err = db.ExecOne("SELECT 1"); err != nil {
		db.Close()
		return nil, nil, err
	}

	/*
		db.AddQueryHook(QueryLogger{
			Logger: logger,
		}) */

	return db, func() {
		db.Close()
	}, nil
}

type DBSlave struct {
	DB *pg.DB
}

func ProvidePostgresSlave(cfg *AppConfig) (dbSlave *DBSlave, cleanup func(), err error) {
	timeoutSecond := cfg.Postgres.ReadTimeout
	if timeoutSecond == 0 {
		timeoutSecond = 60
	}

	dbSlave = new(DBSlave)

	dbSlave.DB = pg.Connect(&pg.Options{
		Addr:        cfg.Postgres.SlaveAddress,
		User:        cfg.Postgres.Username,
		Password:    cfg.Postgres.Password,
		Database:    cfg.Postgres.Database,
		ReadTimeout: time.Duration(timeoutSecond) * time.Second,
	})

	//url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
	//	cfg.Postgres.Username,
	//	cfg.Postgres.Password,
	//	cfg.Postgres.SlaveAddress,
	//	5432,
	//	cfg.Postgres.Database)

	if _, err = dbSlave.DB.ExecOne("SELECT 1"); err != nil {
		dbSlave.DB.Close()
		return nil, nil, err
	}

	/*
		db.AddQueryHook(QueryLogger{
			Logger: logger,
		}) */

	return dbSlave, func() {
		dbSlave.DB.Close()
	}, nil
}
