package app

type Task struct {
	UserID  string
	Message string
}

type User struct {
	TaskDailyLimit int
}
