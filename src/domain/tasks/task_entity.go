package tasks

type (
	Task struct {
		ID          int64  `sql:"primary_key" json:"id"`
		UserId      int64  `sql:"size:100" json:"userId"`
		Content     string `sql:"size:1000" json:"content"`
		CreatedDate string `sql:"default:CURRENT_TIMESTAMP" json:"createdDate"`
	}

	Tasks []Task
)
