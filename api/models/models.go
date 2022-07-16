package models

type APIError struct {
	Message string `json:"message"`
}

type StatusResponse struct {
	Status string `json:"status" validate:"oneof= ok failure"`
}

type TaskCreateRequest struct {
	Content    string `json:"content"`
	TargetDate string `json:"target_date"`
}

type TaskIndexResponse struct {
	Tasks []*Task `json:"tasks"`
}

type UserIndexResponse struct {
	Users []*User `json:"users"`
}

type SettingCreateRequest struct {
	LimitTask int `json:"limit_task"`
}

type SettingUpdateRequest struct {
	LimitTask int `json:"limit_task"`
}

type SettingIndexResponse struct {
	Settings []*Setting `json:"settings"`
}
