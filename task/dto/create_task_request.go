package dto

import (
	"errors"
	apperror "github.com/manabie-com/togo/shared/app_error"
)

var (
	ErrContentNotEmpty = apperror.NewCustomError(errors.New("content is not empty"),
		"content is not empty", "ErrContentNotEmpty")
)

type CreateTaskRequest struct {
	Content string `json:"content"`
}

func (r CreateTaskRequest) Validate() error {
	if r.Content == "" {
		return ErrContentNotEmpty
	}
	return nil
}
