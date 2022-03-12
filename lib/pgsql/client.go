package pgsql

import (
	"fmt"
	"time"

	"github.com/HoangMV/togo/lib/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // import your used driver
)

type SQLClient struct {
	*sqlx.DB
}

func (client *SQLClient) Get() *sqlx.DB {
	return client.DB
}

func NewSqlxDB(c *Config) *SQLClient {
	db, err := sqlx.Connect("postgres", c.DSN)
	if err != nil {
		log.Get().Errorf("NewSqlxDB Connect db error: %+v", err)

		panic(err)
	}

	db.SetMaxOpenConns(c.Active)
	db.SetMaxIdleConns(c.Idle)
	if c.Lifetime < 60 {
		c.Lifetime = 5 * 60
	}
	db.SetConnMaxLifetime(time.Duration(c.Lifetime) * time.Second)

	fmt.Printf("NewSqlxDB: %+v\n", db.Stats())

	return &SQLClient{db}
}
