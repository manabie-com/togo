package entity

type User struct {
	ID    int64
	Name  string
	Todos []Todo
}
