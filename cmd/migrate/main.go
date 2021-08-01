package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/manabie-com/togo/configs"
	"github.com/manabie-com/togo/internal/storages/postgres"
)

var (
	cfg *configs.Config

	dbClient *sql.DB
)

func main() {
	loadConfig()
	if err := loadDatabase(); err != nil {
		panic(err)
	}

	if err := loadMigrations(); err != nil {
		panic(err)
	}
}

func loadConfig() {
	cfg = &configs.Config{
		DBAddress: os.Getenv("DB_ADDRESS"),
	}
}

func loadDatabase() error {
	var err error
	fmt.Println("connect postgres with address", cfg.DBAddress)
	dbClient, err = postgres.NewPostgresClient(cfg.DBAddress)

	return err
}

func loadMigrations() error {
	ctx := context.Background()
	stmt := `CREATE TABLE users (
		id TEXT NOT NULL,
		password TEXT NOT NULL,
		max_todo INTEGER DEFAULT 5 NOT NULL,
		UNIQUE (id),
		CONSTRAINT users_PK PRIMARY KEY (id)
	)`
	if _, err := dbClient.ExecContext(ctx, stmt); err != nil {
		return err
	}
	stmt = `CREATE TABLE tasks (
		id TEXT NOT NULL,
		content TEXT NOT NULL,
		user_id TEXT NOT NULL,
		created_date TEXT NOT NULL,
		number INTEGER DEFAULT 0 NOT NULL,
		UNIQUE (user_id, created_date, number),
		CONSTRAINT tasks_PK PRIMARY KEY (id),
		CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
	)`
	if _, err := dbClient.ExecContext(ctx, stmt); err != nil {
		return err
	}
	return nil
}
