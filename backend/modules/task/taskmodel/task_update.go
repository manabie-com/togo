package taskmodel

type TaskUpdate struct {
	Content *string `json:"content" form:"content"`
	IsDone  *bool   `json:"is_done" form:"is_done"`
}

func (TaskUpdate) TableName() string {
	return Task{}.TableName()
}
