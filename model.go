package togo

import "time"

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	Id        int        `json:"id"`
	Username  string     `json:"username"`
	CreatedAt time.Time  `json:"createdAt"`
	LastLogin *time.Time `json:"lastLogin"`
}

type TaskCreationRequest struct {
	TaskName    string `json:"taskName"`
	Description string `json:"description"`
}

type TaskCreationResponse struct {
	Id          int       `json:"id"`
	TaskName    string    `json:"taskName"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

type TaskLimiterSettingRequest struct {
	UserId int `json:"userId"`
	TaskId int `json:"taskId"`
	Limit  int `json:"limit"`
}

type UserEntity struct {
	Id        int        `db:"id"`
	Username  string     `db:"name"`
	CreatedAt time.Time  `db:"created_at"`
	LastLogin *time.Time `db:"last_login"`
}

type TaskEntity struct {
	Id          int       `db:"id"`
	TaskName    string    `db:"name"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
}

type UserTaskLimitEntity struct {
	UserId    int `db:"user_id"`
	TaskId    int `db:"task_id"`
	Limit     int `db:"num_limit"`
	UpdatedAt int `db:"updated_at"`
}
