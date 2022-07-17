package db

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"github.com/pkg/errors"
)

var (
	// dbManager is a default database manager.
	dbManager *Manager

	// dbMutex prevents dup of default database manager.
	dbMutex sync.Mutex
)

// TransactionalExecutable is a callback function which uses transaction.
type TransactionalExecutable func(context.Context, *sql.Tx) error

// Config contains parameters for database connection.
type Config struct {
	User     string
	Password string
	Host     string
	Database string
}

type Manager struct {
	DB     *sql.DB
	Config Config

	transactionInstanceContextKey      *contextKey
	transactionForceRollbackContextKey *contextKey
}

// transaction contains transaction executor and its identifier.
type transaction struct {
	Transactor *sql.Tx
}

type contextKey struct {
	Name string
}

// Setup itializes default database manager.
func Setup() error {
	c := Config{
		User:     "postgres",
		Password: "password",
		Host:     "postgresql.manabie.todo",
		Database: "todo",
	}

	dbMutex.Lock()
	defer dbMutex.Unlock()

	if dbManager != nil {
		return errors.New("database: already initialized")
	}

	var (
		err error
		dsn = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", c.User, c.Password, c.Host, c.Database)
	)

	DB, err := sql.Open("postgres", dsn)

	if err != nil {
		return errors.Wrap(err, "database: failed to setup database manager")
	}

	DB.SetConnMaxLifetime(10)
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(20)

	if err := DB.Ping(); err != nil {
		return err
	}

	dbManager = &Manager{
		DB:     DB,
		Config: c,

		transactionInstanceContextKey:      &contextKey{"transaction-instance-context-key"},
		transactionForceRollbackContextKey: &contextKey{"transaction-force-rollback-context-key"},
	}

	return nil
}

// Teardown deletes default database manager.
func Teardown() (err error) {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	defer func() {
		if r := recover(); r != nil {
			err = errors.New("database: failed to complete Teardown()")
		}
	}()

	if err := dbManager.DB.Close(); err != nil {
		return errors.Wrap(err, "database: failed to teardown")
	}

	dbManager = nil

	return nil
}

// withTransaction returns a copy of parent context that associated with the new tranasction.
func (dbm *Manager) withTransaction(ctx context.Context, txOpt *sql.TxOptions) (context.Context, error) {
	tx, err := dbm.DB.BeginTx(ctx, txOpt)

	if err != nil {
		return ctx, errors.Wrap(err, "database: failed to begin transaction")
	}

	return context.WithValue(ctx, dbm.transactionInstanceContextKey, transaction{Transactor: tx}), nil
}

func (dbm *Manager) transaction(ctx context.Context, txOpt *sql.TxOptions, fn TransactionalExecutable) (err error) {
	v, hasTransaction := ctx.Value(dbm.transactionInstanceContextKey).(transaction)

	if !hasTransaction || v.Transactor == nil {
		ctx, err = dbm.withTransaction(ctx, txOpt)

		if err != nil {
			return errors.Wrap(err, "database: failed to begin transaction")
		}

		v, _ = ctx.Value(dbm.transactionInstanceContextKey).(transaction)
	}

	tx := v.Transactor

	// CallBack function
	err = fn(ctx, tx)

	if hasTransaction {
		if err != nil {
			return err
		}

		return nil
	}

	if err != nil {
		if rberr := tx.Rollback(); rberr != nil {
			err = errors.Wrapf(err, "rollback errors: %s", rberr.Error())
		}
		return err
	}

	if _, ok := ctx.Value(dbm.transactionForceRollbackContextKey).(struct{}); ok {
		if rberr := tx.Rollback(); rberr != nil {
			err = errors.Wrapf(err, "rollback for testing errors: %s", rberr.Error())
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "database: failed to call tx.Commit()")
	}

	return nil
}

func Transaction(ctx context.Context, txOpt *sql.TxOptions, fn TransactionalExecutable) error {
	if dbManager == nil {
		return errors.New("database: manager is not initialized yet")
	}

	return dbManager.transaction(ctx, txOpt, fn)
}

func TransactionForTesting(ctx context.Context, txOpt *sql.TxOptions, fn TransactionalExecutable) error {
	if dbManager == nil {
		return errors.New("database: manager is not initialized yet")
	}

	// TODO check ENV for testing

	return dbManager.transaction(withForceRollback(ctx, dbManager.transactionForceRollbackContextKey), nil, fn)
}

// withForceRollback returns a copy of parent context that associated with the special flag.
//
// When the context contains this flag, Transaction() always executes tx.Rollback().
func withForceRollback(ctx context.Context, transactionForceRollbackContextKey interface{}) context.Context {
	return context.WithValue(ctx, transactionForceRollbackContextKey, struct{}{})
}
