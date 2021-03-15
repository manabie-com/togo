package tests

import (
	"time"

	"github.com/jmoiron/sqlx"
	d "github.com/manabie-com/togo/internal/todo/domain"
	"github.com/pkg/errors"
)

type TestDB struct {
	dbConn *sqlx.DB
}

func NewTestDB(dbConn *sqlx.DB) *TestDB {
	return &TestDB{dbConn: dbConn}
}

func (t *TestDB) Truncate() error {
	_, err := t.dbConn.Exec("TRUNCATE TABLE users, tasks")
	if err != nil {
		return errors.Wrap(err, "truncate error")
	}

	return nil
}

func (t *TestDB) SeedUser() ([]*d.User, error) {
	users := []*d.User{
		{ID: 1, Username: "test", Password: "testpassword"},
		{ID: 2, Username: "test2", Password: "testpassword2"},
	}

	for i := range users {
		_, err := t.dbConn.Exec("INSERT INTO users (id, username, password) VALUES ($1, $2, crypt($3, gen_salt('bf', 8)))",
			users[i].ID, users[i].Username, users[i].Password)
		if err != nil {
			return nil, errors.Wrap(err, "create user error")
		}
	}

	return users, nil
}

func (t *TestDB) SeedTask() ([]*d.Task, error) {
	now := time.Now()
	time1, _ := time.Parse("2006-01-02", "2021-03-13")
	time2, _ := time.Parse("2006-01-02", "2021-03-15")
	tasks := []*d.Task{
		{ID: 1, Content: "a", UserID: 1, CreatedAt: &time1},
		{ID: 2, Content: "b", UserID: 1, CreatedAt: &time2},
		{ID: 3, Content: "b1", UserID: 1, CreatedAt: &time2},
		{ID: 4, Content: "c", UserID: 1, CreatedAt: &now},
		{ID: 5, Content: "d", UserID: 2, CreatedAt: &now},
	}

	for i := range tasks {
		tk := tasks[i]
		_, err := t.dbConn.Exec("INSERT INTO tasks (id, user_id, content, created_at) VALUES ($1, $2, $3, $4)",
			tk.ID, tk.UserID, tk.Content, tk.CreatedAt.Format("2006-01-02 15:04:05"))
		if err != nil {
			return nil, errors.Wrap(err, "create user error")
		}

	}

	return tasks, nil
}
