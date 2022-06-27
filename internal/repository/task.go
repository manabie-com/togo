package repository

import (
	e "lntvan166/togo/internal/entities"
)

// CREATE

func (r *repository) AddTask(t *e.Task) error {
	const query = `INSERT INTO tasks (
		name, description, created_at, completed, user_id)
		VALUES ($1, $2, $3, $4, $5);`
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, t.Name, t.Description, t.CreatedAt, t.Completed, t.UserID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// READ

func (r *repository) GetAllTask() (*[]e.Task, error) {
	const query = `SELECT * FROM tasks;`
	rows, err := r.DB.Query(query)
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

func (r *repository) GetTaskByID(id int) (*e.Task, error) {
	const query = `SELECT * FROM tasks WHERE id = $1;`
	row := r.DB.QueryRow(query, id)
	var t e.Task
	err := row.Scan(&t.ID, &t.Name, &t.Description, &t.CreatedAt, &t.Completed, &t.UserID)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *repository) GetTasksByUsername(username string) (*[]e.Task, error) {
	const query = `SELECT t.id, t.name, t.description, t.created_at, t.completed
					FROM tasks AS t JOIN users ON t.user_id = users.id
					WHERE users.username = $1;`
	rows, err := r.DB.Query(query, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []e.Task{}

	for rows.Next() {
		var t e.Task
		err := rows.Scan(&t.ID, &t.Name, &t.Description, &t.CreatedAt, &t.Completed)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return &tasks, nil
}

func (r *repository) GetUserIDByTaskID(id int) (int, error) {
	const query = `SELECT user_id FROM tasks WHERE id = $1;`
	row := r.DB.QueryRow(query, id)
	var userID int
	err := row.Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (r *repository) CheckLimitTaskToday(userID int) (bool, error) {
	const query = `SELECT COUNT(*) FROM tasks WHERE user_id = $1 AND created_at >= current_date;`
	row := r.DB.QueryRow(query, userID)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}

	maxTask, err := r.GetMaxTaskByUserID(userID)
	if err != nil {
		return false, err
	}

	if count >= maxTask {
		return true, nil
	}

	return false, nil
}

// UPDATE

func (r *repository) UpdateTask(t *e.Task) error {
	const query = `UPDATE tasks SET name = $1, description = $2, created_at = $3, completed = $4, user_id = $5 WHERE id = $6;`
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, t.Name, t.Description, t.CreatedAt, t.Completed, t.UserID, t.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *repository) CompleteTask(id int) error {
	const query = `UPDATE tasks SET completed = true WHERE id = $1;`
	tx, err := r.DB.Begin()
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

// DELETE

func (r *repository) DeleteTask(id int) error {
	const query = `DELETE FROM tasks WHERE id = $1;`
	tx, err := r.DB.Begin()
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

func (r *repository) DeleteAllTask() error {
	const query = `DELETE FROM tasks;`
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *repository) DeleteAllTaskOfUser(userID int) error {
	const query = `DELETE FROM tasks WHERE user_id = $1;`
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, userID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
