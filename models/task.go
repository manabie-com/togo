package models

import (
	"database/sql"
	"time"
)

type Task struct {
	ID        uint32    `json:"id" validate:"omitempty"`
	Name      string    `json:"name" validate:"required"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UserId    uint32    `json:"userId"`
}

func (t *Task) InsertTask(db *sql.DB) error {
	err := db.QueryRow(`INSERT INTO tasks(name, content, user_id) VALUES($3, $2, $1) RETURNING id, name, content, created_at`, t.UserId, t.Content, t.Name).Scan(&t.ID, &t.Name, &t.Content, &t.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (t *Task) GetTaskByUserId(db *sql.DB) error {
	err := db.QueryRow(`SELECT name, content, created_at FROM tasks WHERE id = $1 AND user_id = $2 `, t.ID, t.UserId).Scan(&t.Name, &t.Content, &t.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (t *Task) GetTasksByUserId(db *sql.DB) ([]*Task, error) {
	var tasks []*Task
	rows, err := db.Query(`SELECT * FROM tasks WHERE user_id = $1`, t.UserId)
	if err != nil {
		return tasks, err
	}

	for rows.Next() {
		var task = &Task{}
		rows.Scan(&task.ID, &task.Name, &task.Content, &task.CreatedAt, &task.UserId)
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (t *Task) IsLimit(db *sql.DB, todayTasksLimit uint) bool {
	// check limit day tasks
	var tasksLength uint
	_ = db.QueryRow(`SELECT COUNT(id)
	FROM tasks 
	WHERE created_at >= NOW() - INTERVAL '24 HOURS' AND user_id = $1`, t.UserId).Scan(&tasksLength)
	return tasksLength == todayTasksLimit
}

func (t *Task) DeleteTaskById(db *sql.DB) error {
	_, err := db.Exec(`DELETE FROM tasks WHERE id = $1 AND user_id = $2`, t.ID, t.UserId)
	if err != nil {
		return err
	}
	return nil
}

func (t *Task) UpdateTaskById(db *sql.DB) error {
	_, err := db.Exec(`UPDATE tasks SET name = $3 WHERE id = $1 AND user_id = $2 RETURNING name, content, created_at`, t.ID, t.UserId, t.Name)
	if err != nil {
		return err
	}
	return nil
}
