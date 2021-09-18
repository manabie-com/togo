package transactional

import (
	"context"
	"database/sql"
	"fmt"
)

func NewDB(sqldb *sql.DB) *db {
	return &db{sqldb}
}

type db struct {
	sqldb *sql.DB
}

func (db *db) Begin(ctx context.Context) (Tx, error) {
	return db.sqldb.BeginTx(ctx, nil)
}

type DB interface {
	Begin(context.Context) (Tx, error)
}

type Tx interface {
	Rollback() error
	Commit() error
}

func withTx(ctx context.Context, db DB, fn func(Tx) error) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction: %v", err)
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()

	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("rollback transaction error: %s (cause of rollback: %w)", rerr, err)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit transaction: %v", err)
	}
	return nil
}

type contextTx struct{}

func NewTxContext(ctx context.Context, tx Tx) context.Context {
	return context.WithValue(ctx, contextTx{}, tx)
}

func TxFromContext(ctx context.Context) Tx {
	tx, _ := ctx.Value(contextTx{}).(Tx)
	return tx
}

func WithTx(ctx context.Context, db DB, fn func(context.Context) error) error {
	if TxFromContext(ctx) != nil {
		// in case tx in tx case we do not begin new transaction
		return fn(ctx)
	}

	return withTx(ctx, db, func(tx Tx) error {
		_ctx := NewTxContext(ctx, tx)
		ctx = withContext(ctx, _ctx)
		return fn(ctx)
	})
}

func withContext(ctx, _ctx context.Context) context.Context {
	return _ctx
}
