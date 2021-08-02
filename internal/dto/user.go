package dto

type CreateUserDTO struct {
	Username string `json:"username" validate:"required,min=6,max=32"`
	Password string `json:"password" validate:"required,min=6,max=32"`
}

type UserLoginDTO struct {
	Username string `json:"username" validate:"required,min=6,max=32"`
	Password string `json:"password" validate:"required,min=6,max=32"`
}
