package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/huuthuan-nguyen/manabie/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

// A TxFn is a function that will be called with an initialized `Transaction` object
// that can be used for executing statements and queries against a database.
type TxFn func(db bun.Tx) error

// WithTransaction creates a new transaction and handles rollback/commit based on the
// error object returned by the `TxFn`
func WithTransaction(ctx context.Context, db *bun.DB, fn TxFn) (err error) {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and re panic
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			_ = tx.Rollback()
		} else {
			// all good, commit
			err = tx.Commit()
		}
	}()

	err = fn(tx)
	return err
}

// NewConnection /**
func NewConnection(c *config.Config) (*bun.DB, error) {
	dns := "postgres://%s:%s@%s:%s/%s?sslmode=disable"

	dnsConnectionString := fmt.Sprintf(dns,
		c.Database.Username,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
	)

	sqlDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dnsConnectionString)))

	return bun.NewDB(sqlDB, pgdialect.New()), nil
}
