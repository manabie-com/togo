package request

import "gopkg.in/go-playground/validator.v9"

type LoginReq struct {
	Username string `json:"username" validate:"required,gte=4,lte=255"`
	Password string `json:"password" validate:"required,gte=6,lte=255"`
}

func Validate(r interface{}) error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}

	return nil
}
