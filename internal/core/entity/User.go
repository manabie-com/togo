package entity

type User struct {
	ID      int64
	Name    string
	MaxTodo int
	Todos   []Todo
}
