package users

// AssignTaskRequest POST users/:id/tasks request
type AssignTaskRequest struct {
	UserID  int16   `json:"-" validate:"min=1"`
	TaskIDs []int16 `json:"task_ids" validate:"min=1,dive,min=1"`
}
