package form

type Form struct {
	UserID     uint   `json:"user_id" form:"user_id" validate:"required,gt=0"`
	TaskDetail string `json:"task_detail" form:"task_detail" validate:"required"`
}
