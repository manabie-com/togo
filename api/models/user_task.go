package models

// Task model represents the task definition
type Task struct {
	Title       string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
}

// User model represents the user definition
type User struct {
	ID       string `bson:"_id" json:"_id"`
	UserName string `bson:"user_name" json:"user_name"`
	MaxTasks int    `bson:"max_tasks" json:"max_tasks"`
}

// UserTask represents the user task definition
type UserTask struct {
	User
	UserID string `bson:"user_id" json:"user_id"`
	Tasks  []Task `bson:"tasks" json:"tasks"`
	InsDay string `bson:"ins_day" json:"ins_day"`
}
