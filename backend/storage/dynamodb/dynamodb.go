package dynamodb

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
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
func NewStorage(dbstring string) (*Storage, error) {
	var cfg, err = config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS config, %w", err)
	}
	client := dynamodb.NewFromConfig(cfg)
	return &Storage{client: client}, nil
}

// GetDBConn returns nil for the dynamodb storage
func (s *Storage) GetDBConn() *sql.DB {
	return nil
}

// IsSQL returns false for the dynamodb storage
func (s *Storage) IsSQL() bool {
	return false
}
