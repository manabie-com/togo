package dynamodb

import (
	"database/sql"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/jssoriao/todo-go/storage"
)

var (
	usersTableName = os.Getenv("USERS_TABLE_NAME")
	tasksTableName = os.Getenv("TASKS_TABLE_NAME")
)

// Storage provides a wrapper around a database and provides
// required methods for interacting with the database
type Storage struct {
	client *dynamodb.Client
}

var _ storage.Storage = (*Storage)(nil)

// NewStorage returns a new dynamodb Storage
func NewStorage(client *dynamodb.Client) *Storage {
	return &Storage{client: client}
}

// GetDBConn returns nil for the dynamodb storage
func (s *Storage) GetDBConn() *sql.DB {
	return nil
}

// IsSQL returns false for the dynamodb storage
func (s *Storage) IsSQL() bool {
	return false
}
