package database

import (
	// "os"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

func CreateConnection() (*pgxpool.Pool, error) {
	connectionString := os.Getenv("POSTGRES_DB_URL")

	dbpool, err := pgxpool.Connect(context.Background(), connectionString)
	
	return dbpool, err
}