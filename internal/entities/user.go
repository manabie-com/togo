package entities

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Plan     string `json:"plan"`
	MaxTodo  int64  `json:"max_todo"`
}

func NewUser() *User {
	return &User{
		Username: "",
		Password: "",
		Plan:     "free",
		MaxTodo:  10,
	}
}
