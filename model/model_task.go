package model

import (
	"errors"
)

type Task struct {
	ID          string `json:"id,omitempty" pg:",notnull"`
	Content     string `pg:"content,type:text" json:"content"`
	UserID      string `pg:"user_id,type:text" json:"user_id"`
	CreatedDate string `pg:"created_date,type:text" json:"created_date"`
}

func (Task) TableName() string {
	return "tasks"
}

func (t *Task) GetTasks(user_id string, created_date string) (task []Task, err error) {
	err = Db.Model(&task).Where(`user_id = ? AND created_date = ?`, user_id, created_date).Select()
	if err != nil {
		return task, err
	}
	return task, nil
}

func (t *Task) InsertTask(task Task) error {
	_, err := Db.Model(&task).Insert()
	if err != nil {
		return err
	}
	return nil
}

func (t *Task) GetListTasks(user_id string) (task []Task, err error) {
	err = Db.Model(&task).Where(`user_id = ?`, user_id).Select()
	if err != nil {
		return task, err
	}
	return task, nil
}

func AddTask(task Task) (*Task, error) {
	var (
		channelUser      = make(chan User)
		channelTask      = make(chan []Task)
		channelErrorUser = make(chan error)
		channelErrorTask = make(chan error)
	)

	go func() {
		var user User
		user, err := user.GetUserById(task.UserID)
		channelUser <- user
		channelErrorUser <- err
	}()

	go func() {
		tasks, err := task.GetTasks(task.UserID, task.CreatedDate)
		channelTask <- tasks
		channelErrorTask <- err
	}()
	user := <-channelUser
	err := <-channelErrorUser
	if err != nil {
		return nil, err
	}
	tasks := <-channelTask
	err = <-channelErrorTask
	if err != nil {
		return nil, err
	}
	defer close(channelUser)
	defer close(channelTask)
	defer close(channelErrorUser)
	defer close(channelErrorTask)

	if len(tasks) >= user.MaxTodo {
		return nil, errors.New("Task today is full")
	}
	err = task.InsertTask(task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func GetListTasks(task Task) (*[]Task, error) {
	var (
		channelTask      = make(chan []Task)
		channelErrorTask = make(chan error)
	)

	go func() {
		tasks, err := task.GetListTasks(task.UserID)
		channelTask <- tasks
		channelErrorTask <- err
	}()

	tasks := <-channelTask
	err := <-channelErrorTask
	if err != nil {
		return nil, err
	}
	defer close(channelTask)
	defer close(channelErrorTask)

	return &tasks, err
}
