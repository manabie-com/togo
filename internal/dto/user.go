package dto

type CreateUserDTO struct {
	Username string `form:"username" validate:"required,min=6,max=32"`
	Password string `form:"password" validate:"required,min=6,max=32"`
}

type UserLoginDTO struct {
	Username string `form:"username" validate:"required,min=6,max=32"`
	Password string `form:"password" validate:"required,min=6,max=32"`
}
