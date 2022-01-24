package create_tasks

import (
	"fmt"
	model3 "togo/module/task/model"
	"togo/module/user/model"
	model2 "togo/module/userconfig/model"
)

var Users = []model.User{
	{
		Id: 1,
		Name: "User Test 1",
		Email: "usertest1@gmail.com",
		Status: 1,
	},
	{
		Id: 2,
		Name: "User Test 2",
		Email: "usertest2@gmail.com",
		Status: 1,
	},
	{
		Id: 3,
		Name: "User Test 3",
		Email: "usertest3@gmail.com",
		Status: 1,
	},
	{
		Id: 4,
		Name: "User Test 4",
		Email: "usertest4@gmail.com",
		Status: 1,
	},
}

var UserConfigs = []model2.UserConfig{
	{
		UserId: 1,
		MaxTask: 4,
	},
	{
		UserId: 2,
		MaxTask: 2,
	},
	{
		UserId: 3,
		MaxTask: 7,
	},
	{
		UserId: 4,
		MaxTask: 9,
	},
}

var UserTasks = []model3.CreateTasksParams{
	{
		UserId: &Users[0].Id,
		Tasks: CreateTasks(),
	},
	{
		UserId: &Users[1].Id,
		Tasks: CreateTasks(),
	},
	{
		UserId: &Users[2].Id,
		Tasks: CreateTasks(),
	},
	{
		UserId: &Users[3].Id,
		Tasks: CreateTasks(),
	},
}

func CreateTasks() []model3.CreateTask {
	data := make([]model3.CreateTask, 0)
	for i := 1; i <= 4; i++ {
		id := uint(i)
		name := fmt.Sprintf("Task %v", i)
		data = append(data, model3.CreateTask{Id: &id, Name: &name})
	}

	return data
}