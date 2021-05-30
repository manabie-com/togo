package tasks

type CreateTasksReq struct {
	Content string `json:"content" form:"content"`
}
