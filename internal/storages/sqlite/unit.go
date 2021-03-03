package sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func CreateUnitTestDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}
	createTables(db)
	return db, nil
}

func createTables(db *sql.DB) {
	const (
		UserId      = "user_id"
		Password    = "password"
		CreatedDate = "2021-03-01"
	)

	_, _ = db.Exec(fmt.Sprintf(`

CREATE TABLE users (
	id TEXT NOT NULL,
	password TEXT NOT NULL,
	max_todo INTEGER DEFAULT 5 NOT NULL,
	CONSTRAINT users_PK PRIMARY KEY (id)
);

INSERT INTO users (id, password) VALUES('%s', '%s');

CREATE TABLE tasks (
	id TEXT NOT NULL,
	content TEXT NOT NULL,
	user_id TEXT NOT NULL,
    created_date TEXT NOT NULL,
	CONSTRAINT tasks_PK PRIMARY KEY (id),
	CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
);

INSERT INTO tasks (id, content, user_id, created_date) VALUES('task_id', 'example_content', '%s', '%s');
	
	`, UserId, Password,
		UserId, CreatedDate))
}
