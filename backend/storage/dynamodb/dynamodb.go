package dynamodb

import (
	"context"
	"database/sql"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/jssoriao/todo-go/storage"
)

var (
	usersTableName = os.Getenv("USERS_TABLE_NAME")
	tasksTableName = os.Getenv("TASKS_TABLE_NAME")
)

type DynamoDBAPI interface {
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
	Query(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error)
	GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
}

type DynamoDBStorage interface {
	CreateUser(storage.User) (storage.User, error)
	GetUser(id string) (*storage.User, error)
	CreateTask(task storage.Task) (storage.Task, error)
	GetTask(userId, id string) (*storage.Task, error)
	CountTasksForTheDay(userID string, dueDate time.Time) (int, error)
}

var _ DynamoDBStorage = (*Storage)(nil)

// Storage provides a wrapper around a database and provides
// required methods for interacting with the database
type Storage struct {
	client DynamoDBAPI
}

var _ storage.Storage = (*Storage)(nil)

// NewStorage returns a new dynamodb Storage
func NewStorage(client DynamoDBAPI) *Storage {
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
