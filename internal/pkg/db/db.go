package db

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type DB struct {
	Conn *pg.DB
}

func (s *DB) Connect() {
	port, err := strconv.Atoi(os.Getenv("POSTGRESQL_PORT"))
	if err != nil {
		logrus.Panicf("postgres port %s is not valid", os.Getenv("POSTGRESQL_PORT"))
	}
	pc := PostgresConfig{
		Host:     os.Getenv("POSTGRESQL_HOST"),
		Port:     uint32(port),
		User:     os.Getenv("POSTGRESQL_USERNAME"),
		Password: os.Getenv("POSTGRESQL_PASSWORD"),
		Database: os.Getenv("POSTGRESQL_DATABASE"),
	}
	connectionString := pc.GetConnectionString()
	opt, err := pg.ParseURL(connectionString)
	if err != nil {
		logrus.Panic("parse go-pg connection string err: %w", err)
	}

	s.Conn = pg.Connect(opt)
	ctx := context.Background()

	if err = s.Conn.Ping(ctx); err != nil {
		panic(err)
	}
	logrus.Info("connect db pg successfully")
}

type PostgresConfig struct {
	Host     string
	Port     uint32
	User     string
	Password string
	Database string
}

func (pc *PostgresConfig) GetConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		pc.User, pc.Password, pc.Host, pc.Port, pc.Database)
}
