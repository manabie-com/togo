// Package store contains create todo related functionality.
package store

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	createtodo "github.com/manabie-com/togo/internal/features/create_todo"
	"github.com/manabie-com/togo/platform/database"
	"go.uber.org/zap"
)

// Store manages the set of APIs for user database access.
type Store struct {
	log    *zap.SugaredLogger
	db     sqlx.ExtContext
	inTran bool
}

// NewStore constructs the api for data access.
func NewStore(log *zap.SugaredLogger, db *sqlx.DB) *Store {
	return &Store{
		log: log,
		db:  db,
	}
}

// WithinTran runs passed function and do commit/rollback at the end.
func (s *Store) WithinTran(ctx context.Context, fn func(s createtodo.Storer) error) error {
	if s.inTran {
		return fn(s)
	}

	f := func(tx *sqlx.Tx) error {
		s := &Store{
			log:    s.log,
			db:     tx,
			inTran: true,
		}
		return fn(s)
	}

	return database.WithinTran(ctx, s.log, s.db.(*sqlx.DB), f)
}

// CreateTodo inserts a new todo into the database.
func (s *Store) CreateTodo(ctx context.Context, todo createtodo.Todo) error {
	const q = `
	INSERT INTO todos
		(id, title, content, user_id, date_created, date_updated)
	VALUES
		(:id, :title, :content, :user_id, :date_created, :date_updated)`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, toDBTodo(todo)); err != nil {
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// GetUserDailyMaxTodo locks the user record and gets their daily maximum number of todos allowed.
func (s *Store) GetUserDailyMaxTodo(ctx context.Context, userID uuid.UUID) (int, error) {
	data := struct {
		ID uuid.UUID `db:"user_id"`
	}{
		ID: userID,
	}

	const q = `
	SELECT daily_max_todo
	FROM users
	WHERE id = :user_id
	FOR UPDATE`

	var maxTodo struct {
		MaxTodo int `db:"daily_max_todo"`
	}

	if err := database.NamedQueryStruct(ctx, s.log, s.db, q, data, &maxTodo); err != nil {
		return 0, fmt.Errorf("namedquerystruct: %w", err)
	}

	return maxTodo.MaxTodo, nil
}

// GetUserTodayTodoCount returns the total number of today todos of a user.
func (s *Store) GetUserTodayTodoCount(ctx context.Context, userID uuid.UUID) (int, error) {
	data := struct {
		ID uuid.UUID `db:"user_id"`
	}{
		ID: userID,
	}

	const q = `
	SELECT count(*)
	FROM todos
	WHERE user_id = :user_id
	AND date(date_created AT TIME ZONE 'Asia/Ho_Chi_Minh') = date(now() AT TIME ZONE 'Asia/Ho_Chi_Minh')`

	var count struct {
		Count int `db:"count"`
	}

	if err := database.NamedQueryStruct(ctx, s.log, s.db, q, data, &count); err != nil {
		return 0, fmt.Errorf("namedquerystruct: %w", err)
	}

	return count.Count, nil
}
