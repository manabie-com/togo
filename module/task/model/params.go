package model

type CreateTasksParams struct {
	Tasks  []CreateTask `json:"tasks"`
	UserId *uint        `json:"user_id"`
}