package model

import (
	"lntvan166/togo/db"
	e "lntvan166/togo/entities"
)

func AddTask(t *e.Task) error {
	const query = `INSERT INTO tasks (
		name, description, created_at, completed, username)
		VALUES ($1, $2, $3, $4, $5);`
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, t.Name, t.Description, t.UserName)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
