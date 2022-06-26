package models

type User struct {
	ID            string `json:"id,omitempty"`
	Username      string `json:"username,omitempty"`
	Password      string `json:"password,omitempty"`
	MaxTaskPerDay int    `json:"max_task_per_day,omitempty"`
}
