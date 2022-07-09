package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lawtrann/togo"
)

type TodoRepo struct {
	DB *DB
}

func NewTodoRepo(db *DB) *TodoRepo {
	return &TodoRepo{DB: db}
}

func (tr *TodoRepo) Add(ctx context.Context, t *togo.Todo, u *togo.User) (*togo.Todo, error) {
	var result togo.Todo

	// Retrieve user by name
	fmt.Println(tr.DB.Now().Format("2006-01-02 15:04:05"), "\n", ISQLTemplate("TodoRepoAddTodo.sql"))
	err := tr.DB.DB.QueryRow(ISQLTemplate("TodoRepoAddTodo.sql"),
		u.ID, t.Description).Scan(&t.ID)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		return nil, err
	}

	result.ID = t.ID
	result.Description = t.Description

	return &result, nil
}

func (tr *TodoRepo) AddWithNewUser(ctx context.Context, t *togo.Todo, u *togo.User) (*togo.Todo, error) {
	var result togo.Todo

	tx, err := tr.DB.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer tx.Commit()

	// Add new user
	fmt.Println(tr.DB.Now().Format("2006-01-02 15:04:05"), "\n", ISQLTemplate("TodoRepoAddNewUser.sql"))
	err = tx.QueryRowContext(ctx, ISQLTemplate("TodoRepoAddNewUser.sql"),
		u.Username, u.LimitedPerDay).Scan(&u.ID)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		tx.Rollback()
		return nil, err
	}

	// Add todo into created user
	fmt.Println(tr.DB.Now().Format("2006-01-02 15:04:05"), "\n", ISQLTemplate("TodoRepoAddTodo.sql"))
	err = tx.QueryRowContext(ctx, ISQLTemplate("TodoRepoAddTodo.sql"),
		u.ID, t.Description).Scan(&t.ID)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		tx.Rollback()
		return nil, err
	}

	result.ID = t.ID
	result.Description = t.Description

	return &result, nil
}
