package request

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/huuthuan-nguyen/manabie/app/model"
	"net/http"
)

type Task struct {
	Content string `json:"content" validate:"required"`
	Status  int    `json:"status" validate:"omitempty"`
}

// Validate /**
func (task *Task) Validate(r *http.Request) error {
	if validate, ok := r.Context().Value("validate").(*validator.Validate); ok {
		return validate.Struct(task)
	}

	return nil
}

// Bind /**
func (task *Task) Bind(r *http.Request, taskModel *model.Task) error {
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(task); err != nil {
		return err
	}

	if err := task.Validate(r); err != nil {
		return err
	}

	taskModel.Content = task.Content
	taskModel.Status = task.Status
	return nil
}
