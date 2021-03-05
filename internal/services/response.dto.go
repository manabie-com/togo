package services

import "github.com/manabie-com/togo/internal/storages"

// response data struct
type TaskListResponse struct {
	Data              []*storages.Task `json:"data"`
	RemainTodayTask   int8             `json:"remain_task_today"`
	MaximumTaskPerDay int8             `json:"maximum_task_perday"`
}
