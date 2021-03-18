package storages

import (
	"database/sql"
	"fmt"

	//"github.com/manabie-com/togo/internal/services"
	"log"
)

// LiteDB for working with sqllite
type LiteDB struct {
	DB *sql.DB
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *LiteDB) RetrieveTasks(userID, createdDate sql.NullString) ([]*Task, error) {
	stmt := "SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2"
	rows, err := l.DB.Query(stmt, userID, createdDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*Task
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
func (l *LiteDB) AddTask(t *Task) error {
	stmt := "INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)"
	_, err := l.DB.Exec(stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}
	return nil
}

// CheckNumTasksInDay return num of tasks match userId added in day
func (l *LiteDB) CheckNumTasksInDay(userId, date sql.NullString) bool {
	stmt := "SELECT\n" +
		"CASE\n" +
		"WHEN t.user_id IS NULL THEN 0\n" +
		"WHEN COUNT(t.user_id) < u.max_todo THEN COUNT(t.user_id)\n" +
		"ELSE NULL\n" +
		"END\n" +
		"FROM tasks t\n" +
		"JOIN users u\n" +
		"ON t.user_id = u.id AND u.id = $1\n" +
		"WHERE t.created_date = $2\n" +
		"GROUP BY t.user_id, u.max_todo"
	row, err := l.DB.Query(stmt, userId, date)
	if err != nil {
		return true
	}
	defer row.Close()

	var count int
	for row.Next() {
		err = row.Scan(&count)
		if err != nil {
			return true
		}
	}
	if err := row.Err(); err != nil {
		return true
	}
	return false
}

// ValidateUser returns tasks if match userID AND password
func (l *LiteDB) ValidateUser(userID, pwd sql.NullString) bool {
	stmt := "SELECT id FROM users WHERE id = $1 AND password = $2"
	row, err := l.DB.Query(stmt, userID, pwd)
	if err != nil {
		return false
	}
	defer row.Close()

	u := &User{}
	for row.Next() {
		err = row.Scan(&u.ID)
		if err != nil {
			return false
		}
	}
	if err = row.Close(); err != nil {
		log.Fatal(err)
	}

	return true
}

func GetConnection(driverName string, host string, port int, user string, password string, dbname string) (*sql.DB, error) {
	datasource := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)
	return sql.Open(driverName, datasource)
}
