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
	stmt := `SELECT id, content, user_id, created_at FROM tasks WHERE user_id = $1`
	if createdDate.String != "" {
		stmt = stmt + ` AND DATE(created_at) = $2`
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
	lastInsertId := 0
	err := l.DB.QueryRow("INSERT INTO tasks (content, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id", t.Content, t.UserID, t.CreatedAt, t.UpdatedAt).Scan(&lastInsertId)
	if err != nil {
		return err
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
		return errUpdate
	}
	t.ID = lastInsertId
	return nil
}

/** ValidateUser check if user existing by query and password hash compare
* @param email, pwd - sql NullString
* @return bool
 */
func (l *DBModel) ValidateUser(email, pwd sql.NullString) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmtEmail := `SELECT password FROM users WHERE email = $1`

	row := l.DB.QueryRowContext(ctx, stmtEmail, email.String)
	var u storages.User
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

/** GetUserFromEmail get user from email
* @param email string
* @return User, error
 */
func (l *DBModel) GetUserFromEmail(email string) (*storages.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `SELECT id, max_todo FROM users WHERE email = $1`
	row := l.DB.QueryRowContext(ctx, stmt, email)
	var u storages.User
	err := row.Scan(&u.ID, &u.MaxTodo)
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
	//check
	stmt := "delete from tasks where id = $1"
	res, err := l.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
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
	stmt := `update tasks set content = $1 where id = $2`
	//ignore the first item and check for an error
	res, err := l.DB.ExecContext(ctx, stmt,
		task.Content,
		task.ID,
	)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	if count <= 0 {
		return errors.New("There's no data to update")
	}
	return nil
}
