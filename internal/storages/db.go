package storages

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type DBModel struct {
	DB *sql.DB
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *DBModel) RetrieveTasks(userID, createdDate sql.NullString) ([]*Task, error) {
	//added timeout for context and cancel if there's something wrong
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `SELECT id, content, user_id, created_at FROM tasks WHERE user_id = $1 AND DATE(created_at) = $2`
	rows, err := l.DB.QueryContext(ctx, stmt, userID.String, createdDate.String)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//initialize array of Task
	var tasks []*Task
	for rows.Next() {
		t := &Task{}
		err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		//push task to the array
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// AddTask adds a new task to DB
func (l *DBModel) AddTask(t Task, u User) (int, error) {
	//added timeout for context and cancel if there's something wrong
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	//get the user ID from the user passed in argument
	userID := u.ID
	var counter int
	now := time.Now()
	dateToday := now.Format("2006-01-02")
	//check the user task today and count
	stmtTask := `select count(id) FROM tasks where user_id = $1 AND DATE(created_at) = $2`
	rowTask := l.DB.QueryRow(stmtTask, userID, dateToday)
	errRowTask := rowTask.Scan(&counter)
	if errRowTask != nil {
		return 0, errRowTask
	}
	//if the users' tasks exceed to 5 this day then throw and error
	if counter >= 5 {
		return 0, errors.New("Only 5 todo task can be created a day")
	}
	//otherwise add task
	lastInsertId := 0
	err := l.DB.QueryRow("INSERT INTO tasks (content, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id", t.Content, t.UserID, t.CreatedAt, t.UpdatedAt).Scan(&lastInsertId)
	if err != nil {
		return 0, err
	}
	//update max todo of the user
	incrementTodo := counter + 1
	stmt := `update users set max_todo = $1, updated_at = $2 where id = $3`
	_, errUpdate := l.DB.ExecContext(ctx, stmt,
		incrementTodo,
		now,
		t.UserID,
	)
	if errUpdate != nil {
		return 0, errUpdate
	}
	return lastInsertId, nil
}

// ValidateUser returns tasks if match email AND password
func (l *DBModel) ValidateUser(email, pwd sql.NullString) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmtEmail := `SELECT password FROM users WHERE email = $1`

	row := l.DB.QueryRowContext(ctx, stmtEmail, email.String)
	var u User
	errRow := row.Scan(&u.Password)
	if errRow != nil {
		return false
	}
	hashedPassword := &u.Password
	password := []byte(*hashedPassword)
	//compare the hash password in the database
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(pwd.String))
	if err != nil {
		return false
	}

	return true
}

//returns one user and error, if any using email query
func (l *DBModel) GetUserFromEmail(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `SELECT id, max_todo FROM users WHERE email = $1`
	row := l.DB.QueryRowContext(ctx, stmt, email)
	var u User
	err := row.Scan(&u.ID, &u.MaxTodo)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
