package models

type Task struct {
	Title       string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
}

type User struct {
	UserName string `bson:"user_name" json:"user_name"`
	MaxTasks int    `bson:"max_tasks" json:"max_tasks"`
	InsDay   string `bson:"ins_day" json:"ins_day"`
	Tasks    []Task `bson:"tasks" json:"tasks"`
}

type UserTask struct {
	User
	Task
}
