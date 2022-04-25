package users

import "time"

// User struct
type User struct {
	UserId      int        `json:"user_id"`
	Name        string     `json:"name"`
	TaskLimit   int        `json:"task_limit"`
	DailyTask   int        `json:"daily_task"`
	LastUpdated time.Time  `json:"last_updated,omitempty"`
	TodoTasks   []TodoTask `json:"todo_tasks"`
}

// TodoTask struct
type TodoTask struct {
	Title      string `json:"title"`
	Detail     string `json:"detail"`
	RemindDate string `json:"remind_date"`
}

// Create a temporary list of user
func CreateTempUsers() []*User {
	users := []*User{
		{
			UserId:    1,
			Name:      "Test User 1",
			TaskLimit: 5,
			DailyTask: 0,
			TodoTasks: []TodoTask{},
		},
		{
			UserId:    2,
			Name:      "Test User 2",
			TaskLimit: 10,
			DailyTask: 0,
			TodoTasks: []TodoTask{},
		},
		{
			UserId:    3,
			Name:      "Test User 3",
			TaskLimit: 20,
			DailyTask: 0,
			TodoTasks: []TodoTask{},
		},
	}
	return users
}
