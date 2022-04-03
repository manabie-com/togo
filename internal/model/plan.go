package model

// Plan represents the plan model
type Plan struct {
	Base
	Name     string `json:"name"`
	MaxTasks int    `json:"max_tasks"`
}
