package models

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"
	"time"
)

type Task struct {
	Id int
	Content string
	Status string
	Time time.Time
	TimeDone time.Time
	UserId int
}
type NewTask struct {
	Content string
	Status string
	Time time.Time
	TimeDone time.Time
	UserId int
}

func GetAllTasks(userId int) ([]Task, error) { // Get all task from the database with user id
	rows, err := DB.Query("SELECT * FROM tasks WHERE userid = $1;", userId)
	var tasks []Task
	if err != nil {
		return tasks, err
	}
	for rows.Next() {
		var task Task
		rows.Scan(&task.Id, &task.Content, &task.Status, &task.Time, &task.TimeDone, &task.UserId)
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func InsertTask(task NewTask) error { // Insert one task to the database
	_, err := DB.Exec("INSERT INTO tasks(content, status,time, timedone, userid) VALUES ($1, $2, $3, $4, $5);", task.Content,task.Status ,task.Time, task.TimeDone, task.UserId)
	if err != nil {
		return errors.New("insert database failed")
	}
	return nil
}

func DeleteTask(id int, userid int) error { // Delete task from database
	_, err := DB.Exec("DELETE FROM tasks WHERE id = $1 AND userid = $2;", id, userid)
	return err
}

func UpdateTask(newTask NewTask, id int, userid int) error { // Update one task already exist in database
	_, err := DB.Exec("UPDATE tasks SET status = COALESCE($1, status) , timedone = COALESCE($2, timedone) WHERE id = $3 AND userid = $4;", newTask.Status, newTask.TimeDone, id, userid)
	if err != nil {
		return err
	}
	return nil
}

func CheckIDTaskAndReturn(w http.ResponseWriter,id int, userId int) (Task, bool) { // Check ID task is valid or not
	task := Task{}
	row := DB.QueryRow("SELECT * FROM tasks WHERE id = $1 AND userid = $2;", id, userId)
	err := row.Scan(&task.Id, &task.Content, &task.Status, &task.Time, &task.TimeDone, &task.UserId)
	if err != nil {
		if err != sql.ErrNoRows {
			http.Error(w, "can't find row", http.StatusFailedDependency)
			return task, false
		}
		return task, false
	}
	return task, true
}

func CheckTaskInput(task NewTask) bool { // Check task input value is valid or not
	Content := strings.TrimSpace(task.Content)
	_ ,validUserid := CheckIDAndReturn(task.UserId)
	if Content == "" || !validUserid{
		return false
	}
	return true
}

