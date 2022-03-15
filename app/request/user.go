package requests

import validation "github.com/go-ozzo/ozzo-validation"

type CreateUserParam struct {
	Name     string `json:"name"`
	Username string `json:"username" `
	Password string `json:"password" `
}

func (param CreateUserParam) Validate() error {
	return validation.ValidateStruct(&param,
		validation.Field(&param.Username, validation.Required),
		validation.Field(&param.Password, validation.Required),
	)
}

type UserSettingPararm struct {
	QuotaPerDay uint `json:"quota_per_day" form:"quota_per_day"`
	UserID      uint `json:"user_id" form:"user_id"`
}

func (param UserSettingPararm) Validate() error {
	return validation.ValidateStruct(&param,
		validation.Field(&param.QuotaPerDay, validation.Required),
		validation.Field(&param.UserID, validation.Required),
	)
}
