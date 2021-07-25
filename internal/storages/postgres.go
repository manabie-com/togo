package storages

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	DB *sql.DB
}

func Initialize(username, password, database string) (Database, error) {
	db := Database{}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, username, password, database)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return db, err
	}
	db.DB = conn
	err = db.DB.Ping()
	if err != nil {
		return db, err
	}
	log.Println("Database connection established")
	return db, nil
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (db Database) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*Task, error) {
	var tasks []*Task
	stmt := `SELECT "id", "content", "user_id", "created_date" FROM "tasks" WHERE "user_id" = $1 AND "created_date" = $2;`
	rows, err := db.DB.Query(stmt, userID, createdDate)
	rows.Scan()
	if err != nil {
		return tasks, err
	}
	defer rows.Close()

	for rows.Next() {
		t := &Task{}
		err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedDate)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// AddTask adds a new task to DB
func (db Database) AddTask(ctx context.Context, t *Task) error {
	stmt := `INSERT INTO "tasks" ("content", "user_id", "created_date") VALUES ($1, $2, $3)`
	_, err := db.DB.ExecContext(ctx, stmt, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

func (db Database) DeleteTaskByDate(ctx context.Context, userId, createdDate sql.NullString) error {
	stmt := `DELETE FROM "tasks" WHERE "user_id" = $1 AND "created_date" = $2`
	_, err := db.DB.ExecContext(ctx, stmt, userId, createdDate)
	if err != nil {
		return err
	}

	return nil
}

// ValidateUser returns tasks if match userID AND password
func (db Database) ValidateUser(ctx context.Context, username, pwd sql.NullString) (*User, error) {
	stmt := `SELECT "id", "user_name", "max_todo" FROM "users" WHERE "user_name" = $1 AND "password" = $2`
	row := db.DB.QueryRowContext(ctx, stmt, username, pwd)
	u := &User{}
	err := row.Scan(&u.ID, &u.Username, &u.MaxTodo)
	return u, err
}

// GetUserById returns user by userId
func (db Database) GetUserById(ctx context.Context, userID sql.NullString) (*User, error) {
	stmt := `SELECT "id", "user_name", "max_todo" FROM "users" WHERE "id" = $1`
	row := db.DB.QueryRowContext(ctx, stmt, userID)
	u := &User{}
	err := row.Scan(&u.ID, &u.Username, &u.MaxTodo)
	return u, err
}

func (db Database) GetUserByUsername(ctx context.Context, username sql.NullString) (*User, error) {
	stmt := `SELECT "id", "user_name", "max_todo" FROM "users" WHERE "user_name" = $1`
	row := db.DB.QueryRowContext(ctx, stmt, username)
	u := &User{}
	err := row.Scan(&u.ID, &u.Username, &u.MaxTodo)
	return u, err
}
