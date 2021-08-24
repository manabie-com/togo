package dtos

type UserDto struct {
}

type TaskDto struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	UserID      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}

type CreateTaskRequest struct {
	Content string `json:"content"`
	UserID  string `json:"-"`
}

type CreateTaskResponse struct {
	Data *TaskDto `json:"data"`
}

type GetListTaskResponse struct {
	Data []*TaskDto `json:"data"`
}

type TokenResponse struct {
	Data string `json:"data"`
}

type ConfigurationDto struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Capacity int64  `json:"capacity"`
	Date     string `json:"date"`
}

type CreateConfigurationRequest struct {
	UserID   string `json:"-"`
	Date     string `json:"date"`
	Capacity int64  `json:"capacity"`
}

type CreateConfigurationResponse struct {
	Data *ConfigurationDto
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewError(err error) *ErrorResponse {
	return &ErrorResponse{Message: err.Error()}
}
