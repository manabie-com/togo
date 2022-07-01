package models

import (
	"database/sql"
	"errors"
	"strings"
	"time"
)

type Task struct {
	Id       int
	Content  string
	Status   string
	Time     time.Time
	TimeDone time.Time
	UserId   int
}
type NewTask struct {
	Content  string
	Status   string
	Time     time.Time
	TimeDone time.Time
	UserId   int
}

// Get all task from the database with user id
func (Conn *DbConn) GetAllTasks(userId int) ([]Task, error) {
	rows, err := Conn.DB.Query(QueryAllTaskText, userId)

	var tasks []Task
	if err != nil {
		return tasks, err
	}
	for rows.Next() {
		var task Task
		if err = rows.Scan(&task.Id, &task.Content, &task.Status, &task.Time, &task.TimeDone, &task.UserId); err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// Insert one task to the database
func (Conn *DbConn) InsertTask(task NewTask) error {
	_, err := Conn.DB.Exec(InsertTaskText, task.Content, task.Status, task.Time, task.TimeDone, task.UserId)
	return err
}

// Delete task from database
func (Conn *DbConn) DeleteTask(id int, userid int) error {
	_, err := Conn.DB.Exec(DeleteTaskText, id, userid)
	return err
}

// Delete task from database
func (Conn *DbConn) DeleteAllTaskFromUser(userid int) error {
	_, err := Conn.DB.Exec(DeleteAllTaskText, userid)
	return err
}

// Update one task already exist in database
func (Conn *DbConn) UpdateTask(newTask Task, id int, userid int) error {
	_, err := Conn.DB.Exec(UpdateTaskText, newTask.Content, newTask.Status, newTask.TimeDone, id, userid)
	return err
}

// Check ID task is valid or not
func (Conn *DbConn) FindTaskByID(id int, userId int) (Task, bool) {
	task := Task{}
	row := Conn.DB.QueryRow(FindTaskByIDText, id, userId)
	err := row.Scan(&task.Id, &task.Content, &task.Status, &task.Time, &task.TimeDone, &task.UserId)
	if err != nil {
		if err != sql.ErrNoRows {
			return task, false
		}
		return task, false
	}
	return task, true
}

// check task per day pass limit of user or not
func (Conn *DbConn) CheckLimitTaskUser(userid int) (bool, error) {
	user, ok := Conn.FindUserByID(userid)
	if !ok {
		return false, errors.New("userid Wrong")
	}
	tasks, err := Conn.GetAllTasks(userid)
	if err != nil || !ok {
		return false, err
	}
	if strings.ToLower(user.Username) == "admin" { // if admin then don't need to check
		return true, nil
	}

	countLimit := 0
	for _, task := range tasks {
		year, month, day := task.Time.Date()
		if year == time.Now().Year() && month == time.Now().Month() && day == time.Now().Day() {
			countLimit++
		}
	}
	return countLimit < user.LimitTask, nil
}
