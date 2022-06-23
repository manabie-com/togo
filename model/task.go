package model

import (
	"lntvan166/togo/db"
	e "lntvan166/togo/entities"
)

func AddTask(t *e.Task) error {
	const query = `INSERT INTO tasks (
		name, description, created_at, completed, user_id)
		VALUES ($1, $2, $3, $4, $5);`
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, t.Name, t.Description, t.UserID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func DeleteTask(id int) error {
	const query = `DELETE FROM tasks WHERE id = $1;`
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func CheckTask(id int) error {
	const query = `UPDATE tasks SET completed = true WHERE id = $1;`
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
