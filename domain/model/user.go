package model

type User struct {
	Id       int
	Username string
	Password string
	Token    string
	Limit    int
}
