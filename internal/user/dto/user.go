package dto

type CreateUserDto struct {
	Name       string `json:"name" validate:"required"`
	LimitCount int    `json:"limit_count" validate:"required,gte=1"`
}
