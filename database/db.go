package database

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

func CreateConnection() (*pgxpool.Pool, error) {
	connectionString := "postgres://todo_user:helloworld@localhost:5432/tododb"

	dbpool, err := pgxpool.Connect(context.Background(), connectionString)
	
	return dbpool, err
}