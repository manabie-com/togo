package model

/*
Handle User model here
*/

type User struct {
	id int
	Limit int
	Task []*Task
}
