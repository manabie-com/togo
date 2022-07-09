package postgres

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

//go:embed integration/*.sql
var integrationFS embed.FS

type DB struct {
	DB     *sql.DB
	Ctx    context.Context // Background context
	Cancel func()          // Cancel background context

	// Datasource name
	DSN string

	// Returns the current time. Defaults to time.Now().
	Now func() time.Time
}

// NewDB returns a new instance of DB associated with the given datasource name.
func NewDB(dsn string) *DB {
	db := &DB{
		DSN: dsn,
		Now: time.Now,
	}
	db.Ctx, db.Cancel = context.WithCancel(context.Background())
	return db
}

// Open opens the database connection.
func (db *DB) Open() (err error) {
	// Ensure a DSN is set before attempting to open the database.
	if db.DSN == "" {
		return fmt.Errorf("DSN required")
	}

	// Connect to the database.
	if db.DB, err = sql.Open("postgres", db.DSN); err != nil {
		return err
	}

	return nil
}

// Close closes the database connection.
func (db *DB) Close() error {
	// Cancel background context.
	db.Cancel()

	// Close database.
	if db.DB != nil {
		return db.DB.Close()
	}
	return nil
}

// Tx wraps the SQL Tx object to provide a timestamp at the start of the transaction.
type Tx struct {
	*sql.Tx
	db  *DB
	now time.Time
}

// BeginTx starts a transaction and returns a wrapper Tx type. This type
// provides a reference to the database and a fixed timestamp at the start of
// the transaction. The timestamp allows us to mock time during tests as well.
func (db *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	tx, err := db.DB.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}

	// Return wrapper Tx that includes the transaction start time.
	return &Tx{
		Tx:  tx,
		db:  db,
		now: db.Now().UTC().Truncate(time.Second),
	}, nil
}

func ISQLTemplate(filename string) string {
	file, err := integrationFS.ReadFile("integration/" + filename)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(file)
}
