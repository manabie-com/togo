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
func (r *Repository)GetAllTasks(userId int) ([]Task, error) { 
	rows, err := r.DB.Query("SELECT * FROM tasks WHERE userid = $1;", userId)
	
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
func (r *Repository)InsertTask(task NewTask) error {
	_, err := r.DB.Exec("INSERT INTO tasks(content, status,time, timedone, userid) VALUES ($1, $2, $3, $4, $5);", task.Content, task.Status, task.Time, task.TimeDone, task.UserId)
	return err
}

// Delete task from database
func (r *Repository)DeleteTask(id int, userid int) error { 
	_, err := r.DB.Exec("DELETE FROM tasks WHERE id = $1 AND userid = $2;", id, userid)
	return err
}

// Delete task from database
func (r *Repository)DeleteAllTaskFromUser(userid int) error { 
	_, err := r.DB.Exec("DELETE FROM tasks WHERE userid = $1;", userid)
	return err
}

 // Update one task already exist in database
func (r *Repository)UpdateTask(newTask Task, id int, userid int) error {
	_, err := r.DB.Exec("UPDATE tasks SET content =COALESCE($1, content), status = COALESCE($2, status), timedone = COALESCE($3, timedone) WHERE id = $4 AND userid = $5;", newTask.Content, newTask.Status, newTask.TimeDone, id, userid)
	return err
}

 // Check ID task is valid or not
func (r *Repository)FindTaskByID(id int, userId int) (Task, bool) {
	task := Task{}
	row := r.DB.QueryRow("SELECT * FROM tasks WHERE id = $1 AND userid = $2;", id, userId)
	err := row.Scan(&task.Id, &task.Content, &task.Status, &task.Time, &task.TimeDone, &task.UserId)
	if err != nil {
		if err != sql.ErrNoRows {
			return task, false
		}
		return task, false
	}
	return task, true
}
func (r *Repository)UpdateStatusAllTask() error {
	_, err := r.DB.Exec("UPDATE tasks SET status = COALESCE($2, status) WHERE time < $1", time.Now())
	return err
}

// Check task input value is valid or not
func CheckTaskInput(task NewTask) bool { 
	var Content string
	if task.Content != ""{
		Content = strings.TrimSpace(task.Content)
	}
	if Content == "" {
		return false
	}
	return true
}

// check task per day pass limit of user or not
func (r *Repository)CheckLimitTaskUser(userid int) (bool, error) { 
	user, ok := r.FindUserByID(userid)
	if !ok {
		return false, errors.New("userid Wrong")
	}
	tasks, err := r.GetAllTasks(userid)
	countLimit :=0
	if strings.ToLower(user.Username) =="admin"{ // if admin then don't need to check
		return true, nil
	}
	if err != nil || !ok{
		return false, err
	}

	for _, task := range tasks{
		year, month, day := task.Time.Date()
		if year == time.Now().Year() && month == time.Now().Month() && day == time.Now().Day(){
			countLimit++;
		}
	}
	return countLimit < user.LimitTask, nil
}