package storage

import (
	"database/sql"
	"time"
)

// Storage contract for storage layer service
type Storage interface {
	// Method that will return database connection for SQL databases or nil for NoSQL databases
	GetDBConn() *sql.DB
	// IsSQL returns true if Storage is an SQL database. Otherwise, returns false.
	IsSQL() bool
}

// User is the persistence model for the user entity
type User struct {
	ID         string    `db:"id" dynamodbav:"id"`
	DailyLimit int       `db:"daily_limit" dynamodbav:"daily_limit"`
	Created    time.Time `db:"created" dynamodbav:"created"`
	Updated    time.Time `db:"updated" dynamodbav:"updated"`
}

// Task is the persistence model for the task entity.
// It uses composite key of (user_id, id)
type Task struct {
	ID      string    `db:"id" dynamodbav:"id"`
	UserID  string    `db:"user_id" dynamodbav:"user_id"`
	Title   string    `db:"title" dynamodbav:"title"`
	Done    bool      `db:"done" dynamodbav:"done"`
	DueDate int64     `db:"due_date" dynamodbav:"due_date"`
	Created time.Time `db:"created" dynamodbav:"created"`
	Updated time.Time `db:"updated" dynamodbav:"updated"`
}
