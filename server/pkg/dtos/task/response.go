package task

type CreateTaskResponse struct {
	TaskID string `json:"id"`
}

type Task struct {
	Id string `json:"id"`
	Content string `json:"content"`
}

type Tasks struct {
	Data []Task `json:"data"`
}

