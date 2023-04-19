package store

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	createtodo "github.com/manabie-com/togo/internal/features/create_todo"
)

// dbTodo represent the structure we need for moving data
// between the app and the database.
type dbTodo struct {
	ID          uuid.UUID      `db:"id"`
	Title       string         `db:"title"`
	Content     sql.NullString `db:"content"`
	UserID      uuid.UUID      `db:"user_id"`
	DateCreated time.Time      `db:"date_created"`
	DateUpdated time.Time      `db:"date_updated"`
}

func toDBTodo(todo createtodo.Todo) dbTodo {
	return dbTodo{
		ID:    todo.ID,
		Title: todo.Title,
		Content: sql.NullString{
			String: todo.Content,
			Valid:  todo.Content != "",
		},
		UserID:      todo.UserID,
		DateCreated: todo.DateCreated,
		DateUpdated: todo.DateUpdated,
	}
}

func toFeatureTodo(dbTd dbTodo) createtodo.Todo {
	usr := createtodo.Todo{
		ID:          dbTd.ID,
		Title:       dbTd.Title,
		Content:     dbTd.Content.String,
		UserID:      dbTd.UserID,
		DateCreated: dbTd.DateCreated.In(time.Local),
		DateUpdated: dbTd.DateUpdated.In(time.Local),
	}

	return usr
}
