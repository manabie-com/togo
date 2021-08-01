package entity

type User struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	MaxTodo  int32  `json:"max_todo"`
}
