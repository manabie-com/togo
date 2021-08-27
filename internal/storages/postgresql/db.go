package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/internal/storages"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Thanhtien1"
	dbname   = "togo"
)

// PostGresDB for working with PostGresql
type PostGresDB struct {
	DB *sql.DB
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *PostGresDB) RetrieveTasks(userID, createdDate sql.NullString) ([]*storages.Task, error) {
	db := OpenConnection()
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2`
	rows, err := db.Query(stmt, userID, createdDate)
	if err != nil {
		return nil, err
	}

	var tasks []*storages.Task
	for rows.Next() {
		t := &storages.Task{}
		err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedDate)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	defer rows.Close()
	defer db.Close()

	return tasks, nil
}

// AddTask adds a new task to DB
func (l *PostGresDB) AddTask(ctx context.Context, t *storages.Task) error {
	db := OpenConnection()
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

// ValidateUser returns tasks if match userID AND password
func (l *PostGresDB) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	stmt := `SELECT id FROM users WHERE id = $1 AND password = $2`
	db := OpenConnection()
	row := db.QueryRowContext(ctx, stmt, userID, pwd)
	u := &storages.User{}
	log.Println(row)
	err := row.Scan(&u.ID)
	if err != nil {
		return false
	}

	return true
}

func (l *PostGresDB) GetMaxTodo(ctx context.Context, userID string) (uint32, error) {
	var maxTodo uint32
	db := OpenConnection()
	stmt := `SELECT max_todo FROM users WHERE id = $1`
	err := db.QueryRowContext(ctx, stmt, userID).Scan(&maxTodo)
	if err != nil {
		return 0, err
	}
	return maxTodo, nil
}

func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
