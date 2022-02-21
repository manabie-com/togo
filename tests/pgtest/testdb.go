// Package dbutils Package test util provides utility func(s) to help integration test.
package dbutils

import (
	"context"
	"fmt"
	"github.com/chi07/todo/db"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"

	"github.com/go-pg/pg/v10"
)

const (
	expirationInSeconds = 120
	dbPoolSize          = 40
	dbPoolTimeout       = 20 * time.Second
	dbPoolMaxWait       = 120 * time.Second
)

// TestUtil represents data to help running integration tests.
type TestUtil struct {
	DB         *pg.DB
	DBURL      string
	PostgresDB *sqlx.DB

	Log     *logrus.Entry
	Context context.Context
}

// New creates a pointer to the instance of test-util
func New() *TestUtil {
	log := logrus.New()
	return &TestUtil{
		Log:     logrus.NewEntry(log),
		Context: context.Background(),
	}
}

func (util *TestUtil) bootstrapDB() error {
	pool, err := dockertest.NewPool("")
	if err != nil {
		return err
	}

	// pulls an image, creates a container based on it and runs
	res, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "13",
		Env: []string{
			"POSTGRES_USER=postgres",
			"POSTGRES_PASSWORD=postgres",
			"POSTGRES_DB=todos",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		return err
	}

	if err := res.Expire(expirationInSeconds); err != nil {
		return err
	}

	if err := util.connectDB(pool, res); err != nil {
		return err
	}

	return nil
}

func (util *TestUtil) connectDB(pool *dockertest.Pool, containerResource *dockertest.Resource) error {
	address := containerResource.GetHostPort("5432/tcp")
	dbURL := fmt.Sprintf("postgres://postgres:postgres@%s/todos?sslmode=disable", address)
	util.DBURL = dbURL
	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	pool.MaxWait = dbPoolMaxWait
	return pool.Retry(func() error {
		dbOptions, err := pg.ParseURL(dbURL)
		dbOptions.PoolSize = dbPoolSize
		dbOptions.PoolTimeout = dbPoolTimeout
		if err != nil {
			return err
		}

		util.DB = pg.Connect(dbOptions)
		if err := util.DB.Ping(util.Context); err != nil {
			return err
		}

		db, err := sqlx.Connect("postgres", dbURL)
		if err != nil {
			logrus.Fatal("could not connect to db")
		}
		util.PostgresDB = db

		return nil
	})
}

// InitDB initializes db.
func (util *TestUtil) InitDB() error {
	if err := util.bootstrapDB(); err != nil {
		return err
	}

	return nil
}

// SetupDB do creating schemas and populate data to database
func (util *TestUtil) SetupDB() error {
	if err := db.Migrate("file://../../db/migrations", util.DBURL); err != nil {
		return err
	}
	if err := util.doSeeding(); err != nil {
		return err
	}

	return nil
}

func (util *TestUtil) doSeeding() error {
	return nil
}

// CleanAndClose cleanup db and closes connection.
func (util *TestUtil) CleanAndClose() {
	// Drop entire schema, the migrator will recreate it
	//if _, err := util.DB.Exec(`drop schema public cascade;`); err != nil {
	//	util.Log.Infof("query execution error %v", err)
	//}
}
