package repository

import (
	"database/sql/driver"
	"errors"
	"github.com/stretchr/testify/require"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"manabieAssignment/internal/core/entity"
	"regexp"
	"testing"
	"time"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestTodoRepository_CreateTodo_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	require.NoError(t, err)
	todo := entity.Todo{
		UserID:  1,
		Name:    "todo_name",
		Content: "todo_content",
	}
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "todos" ("created_at","updated_at","deleted_at","user_id","name","content") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
		WithArgs(AnyTime{}, AnyTime{}, nil, todo.UserID, todo.Name, todo.Content).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1)).
		WillReturnError(nil)
	todoRepo := NewTodoRepository(gormDB)
	todoId, err := todoRepo.CreateTodo(todo)
	require.NoError(t, err)
	require.Equal(t, uint(1), todoId)
}

func TestTodoRepository_CreateTodo_Failed(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	require.NoError(t, err)
	todo := entity.Todo{
		UserID:  1,
		Name:    "todo_name",
		Content: "todo_content",
	}
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "todos" ("created_at","updated_at","deleted_at","user_id","name","content") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
		WithArgs(AnyTime{}, AnyTime{}, nil, todo.UserID, todo.Name, todo.Content).
		WillReturnError(errors.New("something wrong"))
	todoRepo := NewTodoRepository(gormDB)
	_, err = todoRepo.CreateTodo(todo)
	require.Error(t, err)
}
