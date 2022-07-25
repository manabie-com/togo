package request

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/huuthuan-nguyen/manabie/app/model"
	"net/http"
)

type Credentials struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Validate /**
func (credentials *Credentials) Validate(r *http.Request) error {
	if validate, ok := r.Context().Value("validate").(*validator.Validate); ok {
		return validate.Struct(credentials)
	}

	return nil
}

// Bind /**
func (credentials *Credentials) Bind(r *http.Request, userModel *model.User) error {
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(credentials); err != nil {
		return err
	}

	if err := credentials.Validate(r); err != nil {
		return err
	}

	userModel.Email = credentials.Email
	userModel.Password = credentials.Password
	return nil
}
