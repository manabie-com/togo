package psql

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/manabie-com/togo/internal/storages"
	"golang.org/x/crypto/bcrypt"
)

type DBModel struct {
	DB *sql.DB
}

//NewModels returns models with db pool
func NewModels(db *sql.DB) *DBModel {
	return &DBModel{
		DB: db,
	}
}

const (
	sqlValidateUser     = `SELECT password, email FROM users WHERE email = ?`
	sqlAddTask          = `INSERT INTO tasks (content, user_id, created_at, updated_at) VALUES (?, ?, ?, ?)`
	sqlRetrieveTasks    = `SELECT id, content, user_id, created_at FROM tasks WHERE user_id = ? AND DATE(created_at) = ?`
	sqlGetUserFromEmail = `SELECT id, max_todo, email FROM users WHERE email = ?`
	sqlUpdateTask       = `update tasks set content = $1 where id = $2`
	sqlDeleteTask       = `delete from tasks where id = ?`
)

/** RetrieveTasks returns tasks if match userID AND createDate.
* @param email, createdDate - sql NullString
* @return Task, error
 */
func (l *DBModel) RetrieveTasks(email, createdDate sql.NullString) ([]*storages.Task, error) {
	//added timeout for context and cancel if there's something wrong
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var user storages.User
	m, err := l.GetUserFromEmail(email.String)
	if err != nil {
		return nil, err
	}
	user = *m
	userID := user.ID
	//initialize variables to be assigned on conditions if there will be a date query
	var rowsDB *sql.Rows
	var errDB error
	stmt := `SELECT id, content, user_id, created_at FROM tasks WHERE user_id = ?`
	if createdDate.String != "" {
		stmt = stmt + ` AND DATE(created_at) = ?`
		rows, err := l.DB.QueryContext(ctx, stmt, userID, createdDate.String)
		rowsDB = rows
		errDB = err
	} else {
		rows, err := l.DB.QueryContext(ctx, stmt, userID)
		rowsDB = rows
		errDB = err
	}
	if errDB != nil {
		return nil, errDB
	}
	defer rowsDB.Close()
	//initialize array of Task
	var tasks []*storages.Task
	for rowsDB.Next() {
		t := &storages.Task{}
		err := rowsDB.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		//push task to the array
		tasks = append(tasks, t)
	}

	if err := rowsDB.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

/** AddTask adds a new task to DB
* @param Task, email string
* @return int, error
 */
func (l *DBModel) AddTask(t *storages.Task, email string) error {
	//added timeout for context and cancel if there's something wrong

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var user storages.User
	m, errGetUser := l.GetUserFromEmail(email)
	if errGetUser != nil {
		return errGetUser
	}
	user = *m
	//get the user ID from the user passed in argument
	userID := user.ID
	t.UserID = userID
	var counter int
	now := time.Now()
	dateToday := now.Format("2006-01-02")
	//check the user task today and count
	stmtTask := `select count(id) FROM tasks where user_id = $1 AND DATE(created_at) = $2`
	rowTask := l.DB.QueryRow(stmtTask, userID, dateToday)
	errRowTask := rowTask.Scan(&counter)
	if errRowTask != nil {
		return errRowTask
	}
	//if the users' tasks exceed to 5 this day then throw and error
	if counter >= 5 {
		return errors.New("Only 5 todo task can be created a day")
	}
	//otherwise add task
	tx, errDB := l.DB.Begin()
	if errDB != nil {
		return errDB
	}

	_, err := l.DB.ExecContext(ctx, sqlAddTask, &t.Content, &t.UserID, t.CreatedAt, t.UpdatedAt)
	if err != nil {
		return err
	}
	defer func() {
		switch err {
		case nil:
			_ = tx.Commit()
		default:
			_ = tx.Rollback()
		}
	}()
	//update max todo of the user
	incrementTodo := counter + 1
	stmt := `update users set max_todo = ?, updated_at = ? where id = ?`
	_, errUpdate := l.DB.ExecContext(ctx, stmt,
		incrementTodo,
		t.UpdatedAt,
		t.UserID,
	)

	defer func() {
		switch errUpdate {
		case nil:
			_ = tx.Commit()
		default:
			_ = tx.Rollback()
		}
	}()
	if errUpdate != nil {
		return errUpdate
	}
	return nil
}

/** ValidateUser check if user existing by query and password hash compare
* @param email, pwd - sql NullString
* @return bool
 */
func (l *DBModel) ValidateUser(email, pwd sql.NullString) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := l.DB.QueryRowContext(ctx, sqlValidateUser, email.String)
	var u storages.User
	errRow := row.Scan(&u.Password, &u.Email)
	if errRow != nil {
		return false
	}
	hashedPassword := u.Password
	password := []byte(hashedPassword)
	//compare the hash password in the database
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(pwd.String))
	if err != nil {
		return false
	}

	return true
}

/** GetUserFromEmail get user from email
* @param email string
* @return User, error
 */
func (l *DBModel) GetUserFromEmail(email string) (*storages.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := l.DB.QueryRowContext(ctx, sqlGetUserFromEmail, email)
	var u storages.User
	err := row.Scan(&u.ID, &u.MaxTodo, &u.Email)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

/**Deletion of task
 * @param id int
 * @return error
 */
func (l *DBModel) DeleteTask(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	tx, errDB := l.DB.Begin()
	if errDB != nil {
		return errDB
	}
	res, err := l.DB.ExecContext(ctx, sqlDeleteTask, id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	defer func() {
		switch err {
		case nil:
			_ = tx.Commit()
		default:
			_ = tx.Rollback()
		}
	}()
	if err != nil {
		panic(err)
	}
	if count <= 0 {
		return errors.New("There's no existing data to delete")
	}

	return nil
}

/**UpdateTask of task
 * @param task Task
 * @return error
 */
func (l *DBModel) UpdateTask(task *storages.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	tx, errDB := l.DB.Begin()
	if errDB != nil {
		return errDB
	}
	//ignore the first item and check for an error
	res, err := l.DB.ExecContext(ctx, sqlUpdateTask,
		task.Content,
		task.ID,
	)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	defer func() {
		switch err {
		case nil:
			_ = tx.Commit()
		default:
			_ = tx.Rollback()
		}
	}()
	if err != nil {
		panic(err)
	}
	if count <= 0 {
		return errors.New("There's no data to update")
	}
	return nil
}
