package http

import "github.com/datshiro/togo-manabie/internal/interfaces/models"

func NewUserResponse(user *models.User) UserResponse {
	return UserResponse{
		Name:  user.Name,
		Email: user.Email.String,
		Quota: user.Quota,
		Tasks: NewUserTaskResponse(user.R.Tasks),
	}
}

func NewUserTaskResponse(tasks models.TaskSlice) (resp []UserTaskResponse) {
	for _, task := range tasks {
		resp = append(resp, UserTaskResponse{
			Title:       task.Title,
			Description: task.Description.String,
			Priority:    task.Priority,
			IsDone:      task.IsDone,
		})
	}
	return resp
}

type UserResponse struct {
	Name  string             `json:"name"`
	Email string             `json:"email"`
	Quota int                `json:"quota"`
	Tasks []UserTaskResponse `json:"tasks"`
}

type UserTaskResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	IsDone      bool   `json:"is_done"`
}
