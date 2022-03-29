package rest

type Task struct {
	ID      string `json:"id,omitempty"`
	UserID  string `json:"userId,omitempty"`
	Message string `json:"message,omitempty"`
}

type User struct {
	ID             string `json:"id,omitempty"`
	TaskDailyLimit int    `json:"taskDailyLimit,omitempty"`
}
