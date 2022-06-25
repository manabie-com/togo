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

func (t *Task) InsertOne(db *sql.DB) error {
	err := db.QueryRow(`INSERT INTO tasks(name, content, user_id) VALUES($3, $2, $1) RETURNING name, content, created_at`, t.UserId, t.Content, t.Name).Scan(&t.Name, &t.Content, &t.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (t *Task) GetOneByUserId(db *sql.DB) error {
	err := db.QueryRow(`UPDATE tasks SET name = $3 WHERE id = $1 AND user_id = $2 RETURNING name, content, created_at`, t.ID, t.UserId, t.Name).Scan(&t.Name, &t.Content, &t.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
