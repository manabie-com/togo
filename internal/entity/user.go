package entity

type User struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	MaxTodo  int32  `json:"max_todo"`
}
