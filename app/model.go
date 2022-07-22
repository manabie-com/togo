package app

import (
	"database/sql"
	"time"
)

type Todo struct {
	Id           int       `json:"id"`
	Content      string    `json:"content"`
	User_id      int       `json:"user_id"`
	Created_date time.Time `json:"created_date"`
}

func (todo *Todo) createTodo(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO todos (content, user_id, created_date) VALUES ($1, $2, $3) RETURNING id",
		todo.Content, todo.User_id, time.Now(),
	).Scan(&todo.Id)

	if err != nil {
		return err
	}

	return nil
}
