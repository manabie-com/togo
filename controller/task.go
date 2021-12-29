package controller

import (
	"fmt"
	"strconv"
	"time"
	"togo/db"
	"togo/model"
)


func GetAllTaskByUser(userId int,createdDate string) ([]*model.Task,error) {
	var task []*model.Task

	if err := db.DB.Where("user_id = ? AND created_date =?",strconv.Itoa(userId),createdDate).Find(&task).Error; err != nil {
		return nil,fmt.Errorf("Record not found")
	}

	return task,nil
}

func AddTask(user model.User,content string) (*model.Task,error) {
	current := time.Now()

	task := &model.Task{
		Content:     content,
		UserID:      user.Id,
		CreatedDate: current.Format("2006-01-02"),
	}

	listTask,err := GetAllTaskByUser(user.Id,task.CreatedDate)
	if err != nil {
		return nil,err
	}

	if len(listTask) >=user.MaxTodo {
		return nil,fmt.Errorf("user reach task limit error")
	}

	return task,nil
}

func CreateTask(task *model.Task)  {
	db.DB.Create(task)
}
