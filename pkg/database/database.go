package database

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
)

type Connection interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type CommandFunc func(context.Context, Connection) error

type Database interface {
	Connect(cfg *Config) error
	Disconnect() error
	WithoutTransaction(ctx context.Context, cmdFuncs ...CommandFunc) error
	Transaction(ctx context.Context, cmdFuncs ...CommandFunc) error
}

func NewDatabase(connector Connector) Database {
	return &database{
		connector: connector,
	}
}

type database struct {
	connector Connector
}

func (p *database) Connect(cfg *Config) error {
	return p.connector.Open(cfg)
}

func (p *database) Disconnect() error {
	return p.connector.Close()
}

func (p *database) WithoutTransaction(ctx context.Context, cmdFuncs ...CommandFunc) error {
	db := p.connector.GetDB()
	if db == nil {
		return errors.New("forgot connect to database")
	}
	var err error
	for _, cmdFunc := range cmdFuncs {
		err = cmdFunc(ctx, db)
		if err != nil && !reflect.ValueOf(err).IsNil() {
			return err
		}
	}
	return nil
}

func (p *database) Transaction(ctx context.Context, cmdFuncs ...CommandFunc) error {
	db := p.connector.GetDB()
	if db == nil {
		return errors.New("forgot connect to database")
	}

	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	for _, cmdFunc := range cmdFuncs {
		err = cmdFunc(ctx, tx)
		if err != nil && !reflect.ValueOf(err).IsNil() {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}
