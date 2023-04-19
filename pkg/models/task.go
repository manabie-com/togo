package models

type Task struct {
	TaskId      string `json:"task_id"`
	UserId      int64  `json:"user_id"`
	Content     string `json:"content"`
	CreatedDate string `json:"created_date"`
	EventTime   string `json:"event_time"`
}

type CountTask struct {
	NumberTask int64 `json:"number_task"`
	LimitTask  int64 `json:"limit_task"`
}
