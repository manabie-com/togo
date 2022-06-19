package postgres

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"testing"
	"togo/internal/pkg/db"
)

var postgresDB *db.DB
var userRepo *UserRepository
var taskRepo *TaskRepository

const connectionString = "postgres://postgres:postgres@localhost:5432/togo?sslmode=disable"

func TestMain(m *testing.M) {
	logrus.Info("connecting to postgres")
	postgresDB = &db.DB{}
	opt, err := pg.ParseURL(connectionString)
	if err != nil {
		logrus.Panic("parse go-pg connection string err: %w", err)
	}

	postgresDB.Conn = pg.Connect(opt)
	ctx := context.Background()

	if err = postgresDB.Conn.Ping(ctx); err != nil {
		panic(err)
	}
	userRepo = &UserRepository{DB: postgresDB}
	taskRepo = &TaskRepository{DB: postgresDB}
	m.Run()
}
