package request

type SignUp struct {
	Email    string `validate:"required,email,max=300"`
	Password string `validate:"required,min=8"`
	MaxTodo  int64  `json:"max_todo" validate:"required,min=1"`
}

type SignIn struct {
	Email    string `validate:"required,email,max=300"`
	Password string `validate:"required,min=8"`
}
