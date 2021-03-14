package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	c "github.com/manabie-com/togo/internal/pkg/config"
)

const (
	CONN_MAX_LIFETIME = time.Second * 30
	MAX_IDLE_CONNS    = 500
	MAX_OPEN_CONNS    = 250
)

var db *sqlx.DB

func GetConnection() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	timezone := os.Getenv("TZ")
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&timezone=%s",
		dbUser, dbPass, dbHost, dbPort, dbName, timezone)
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	maxAttempt := c.GetEnvInt("DB_NUM_ATTEMPT")
	if err := WaitingForDB(conn, maxAttempt); err != nil {
		return nil, err
	}

	return conn, nil
}

func NewSQLXConn() *sqlx.DB {
	if db == nil {
		conn, err := GetConnection()
		if err != nil {
			panic(err)
		}

		db = sqlx.NewDb(conn, "postgres")
		db.DB.SetConnMaxLifetime(CONN_MAX_LIFETIME)
		db.DB.SetMaxIdleConns(MAX_IDLE_CONNS)
		db.DB.SetMaxOpenConns(MAX_OPEN_CONNS)
	}

	return db
}
