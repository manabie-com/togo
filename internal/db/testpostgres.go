package db

import (
	"database/sql"
	"time"

	"github.com/pkg/errors"
	"github.com/manabie-com/togo/internal/storages"
)

func Truncate(db *sql.DB) error {
	stmt := "TRUNCATE TABLE tasks, users;"

	if _, err := db.Exec(stmt); err != nil {
		return errors.Wrap(err, "truncate test database tables")
	}

	return nil
}

func SeedUsers(db *sql.DB) ([]storages.User, error) {
	users := []storages.User{
		{
			ID:     "firstUser",
			Password:  "example",
		},
		{
			ID:     "secondUser",
			Password:  "example",
		},
	}

	for i := range users {
		stmt, err := db.Prepare("INSERT INTO users (ID, password) VALUES ($1, $2);")
		if err != nil {
			return nil, errors.Wrap(err, "prepare users insertion")
		}

		stmt.QueryRow(users[i].ID, users[i].Password)

		if err := stmt.Close(); err != nil {
			return nil, errors.Wrap(err, "close psql statement")
		}
	}

	return users, nil
}

func SeedTasks(db *sql.DB, users []storages.User) ([]storages.Task, error) {
	tasks := []storages.Task{
		{
			ID:     "2d0a8118-dc77-4187-8cbe-de187425ce38",
			Content:  "content 1",
			UserID: users[0].ID,
			CreatedDate: time.Now().Format("2006-01-02"),
		},
		{
			ID:     "24a3d2b5-b358-48b6-bdad-340ac52a2d54",
			Content:  "content 2",
			UserID: users[1].ID,
			CreatedDate: time.Now().Format("2006-01-02"),
		},
	}

	for i := range tasks {
		stmt, err := db.Prepare("INSERT INTO tasks (ID, content, user_id, created_date) VALUES ($1, $2, $3, $4);")
		if err != nil {
			return nil, errors.Wrap(err, "prepare tasks insertion")
		}

		stmt.QueryRow(tasks[i].ID, tasks[i].Content, tasks[i].UserID, tasks[i].CreatedDate)

		if err := stmt.Close(); err != nil {
			return nil, errors.Wrap(err, "close psql statement")
		}
	}

	return tasks, nil
}
