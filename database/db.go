package database

import (
	// "os"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

func CreateConnection() (*pgxpool.Pool, error) {
	// connectionString := os.Getenv("POSTGRES_DB_URL")
	connectionString := "postgres://todo_user:secret@localhost:5432/tododb"

	dbpool, err := pgxpool.Connect(context.Background(), connectionString)
	
	return dbpool, err
}