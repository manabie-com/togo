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

func GetAllTask() (*[]e.Task, error) {
	const query = `SELECT * FROM tasks;`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []e.Task{}

	for rows.Next() {
		var t e.Task
		err := rows.Scan(&t.ID, &t.Name, &t.Description, &t.CreatedAt, &t.Completed, &t.UserID)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return &tasks, nil
}

func GetTaskByUserID(userID int) (*[]e.Task, error) {
	const query = `SELECT * FROM tasks WHERE user_id = $1;`
	rows, err := db.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []e.Task{}

	for rows.Next() {
		var t e.Task
		err := rows.Scan(&t.ID, &t.Name, &t.Description, &t.CreatedAt, &t.Completed, &t.UserID)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return &tasks, nil
}
