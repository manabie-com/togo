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
}

// Setup itializes default database manager.
func Setup(ctx context.Context, c Config) error {
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

	if err := DB.Ping(); err != nil {
		return err
	}

	dbManager = &Manager{
		DB:     DB,
		Config: c,
	}

	return nil
}

// Teardown deletes default database manager.
func (dbm *Manager) Teardown() (err error) {
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

func (dbm *Manager) transaction(ctx context.Context, txOpt *sql.TxOptions, fn TransactionalExecutable) error {
	tx, err := dbm.DB.BeginTx(ctx, txOpt)
	if err != nil {
		return errors.Wrap(err, "database: failed to begin transaction")
	}

	// CallBack function
	err = fn(ctx, tx)

	if err != nil {
		if rberr := tx.Rollback(); rberr != nil {
			err = errors.Wrapf(err, "rollback errors: %s", rberr.Error())
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
