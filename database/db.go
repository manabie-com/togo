package database

import (
	"os"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Create the database connection to POSTGRESQL DB where POSTGRES_DB_URL is an OS environment variable 
func CreateConnection() (*pgxpool.Pool, error) {
	connectionString := os.Getenv("POSTGRES_DB_URL")

	dbpool, err := pgxpool.Connect(context.Background(), connectionString)
	
	return dbpool, err
}