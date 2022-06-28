package repository

import (
	"github.com/stretchr/testify/require"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"manabieAssignment/internal/core/entity"
	"regexp"
	"testing"
	"time"
)

func TestTodoRepository_CreateTodo_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})
	require.NoError(t, err)
	todo := entity.Todo{
		UserID:    1,
		Name:      "todo_name",
		Content:   "todo_content",
		CreatedAt: time.Now(),
	}
	mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO todos (created_at,updated_at,deleted_at,user_id,name,content) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id")).
		ExpectQuery().WithArgs(todo.CreatedAt, time.Now(), nil, todo.UserID, todo.Name, todo.Content).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	todoRepo := NewTodoRepository(gormDB)
	todoId, err := todoRepo.CreateTodo(todo)
	require.NoError(t, err)
	require.Equal(t, todoId, 1)
}
