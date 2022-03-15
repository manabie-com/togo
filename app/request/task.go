package requests

import validation "github.com/go-ozzo/ozzo-validation"

type CreateTaskParam struct {
	Name        string `json:"name"`
	Description string `json:"description" `
}

func (param CreateTaskParam) Validate() error {
	return validation.ValidateStruct(&param,
		validation.Field(&param.Name, validation.Required),
		validation.Field(&param.Description, validation.Required),
	)
}

type PagineUserTaskParam struct {
	Limit int `json:"limit" form:"limit"`
	Page  int `json:"page" form:"page"`
}

type UpdateTaskParam struct {
	Name        *string `json:"name"`
	Description *string `json:"description" `
	TaskID      int     `json:"task_id" `
}

func (param UpdateTaskParam) Validate() error {
	return validation.ValidateStruct(&param,
		validation.Field(&param.TaskID, validation.Required),
	)
}
