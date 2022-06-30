package repository

import (
	"database/sql"
	e "lntvan166/togo/internal/entities"
)

type taskRepository struct {
	DB *sql.DB
}

func NewTaskRepository(db *sql.DB) *taskRepository {
	return &taskRepository{
		DB: db,
	}
}

// CREATE

func (r *taskRepository) CreateTask(t *e.Task) (int, error) {
	const query = `INSERT INTO tasks (
		name, description, created_at, completed, user_id)
		VALUES ($1, $2, $3, $4, $5) RETURNING id;`
	tx, err := r.DB.Begin()
	if err != nil {
		return 0, err
	}
	row := tx.QueryRow(query, t.Name, t.Description, t.CreatedAt, t.Completed, t.UserID)
	err = row.Scan(&t.ID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()
	return t.ID, nil
}

// READ

func (r *taskRepository) GetAllTask() (*[]e.Task, error) {
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

func (r *taskRepository) GetTaskByID(id int) (*e.Task, error) {
	const query = `SELECT * FROM tasks WHERE id = $1;`
	row := r.DB.QueryRow(query, id)
	var t e.Task
	err := row.Scan(&t.ID, &t.Name, &t.Description, &t.CreatedAt, &t.Completed, &t.UserID)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *taskRepository) GetTasksByUserID(id int) (*[]e.Task, error) {
	const query = `SELECT * FROM tasks WHERE user_id = $1;`
	rows, err := r.DB.Query(query, id)
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

func (r *taskRepository) GetNumberOfTaskTodayByUserID(id int) (int, error) {
	const query = `SELECT COUNT(*) FROM tasks WHERE user_id = $1 AND DATE(created_at) = CURRENT_DATE`
	var count int
	err := r.DB.QueryRow(query, id).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// func (r *taskRepository) GetMaxTaskByUserID(id int) (int, error) {
// 	const query = `SELECT max_todo FROM users WHERE id = $1`
// 	var max int
// 	err := r.DB.QueryRow(query, id).Scan(&max)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return max, nil
// }

// UPDATE

func (r *taskRepository) UpdateTask(t *e.Task) error {
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

func (r *taskRepository) CompleteTask(id int) error {
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

func (r *taskRepository) DeleteTask(id int) error {
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

func (r *taskRepository) DeleteAllTaskOfUser(userID int) error {
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

// for testing
func (r *taskRepository) RollbackFromDelete(task *e.Task) error {
	const query = `INSERT INTO tasks (
		id, name, description, created_at, completed, user_id)
		VALUES ($1, $2, $3, $4, $5, $6);`
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, task.ID, task.Name, task.Description, task.CreatedAt, task.Completed, task.UserID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
