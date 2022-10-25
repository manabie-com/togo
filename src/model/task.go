package model

/*
Handle Task model here
*/

type Task struct {
	id int
	Name string
	User []*User
}
