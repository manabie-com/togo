package tasks

type (
	Task struct {
		ID          int64  `sql:"primary_key" json:"id"`
		UserID      int64  `sql:"user_id" json:"userId"`
		Content     string `sql:"content" json:"content"`
		CreatedDate string `sql:"created_date,default:CURRENT_TIMESTAMP" json:"createdDate"`
	}

	Tasks []Task
)
