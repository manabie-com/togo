package models

type User struct {
	Id int 
	Name string
	MaxNumberOfTasks int
}

func MakeUser(
	iId int,
	iName string,
	iMaxNumberOfTasks int,
) User {
	return User{
		Id: iId,
		Name: iName,
		MaxNumberOfTasks: iMaxNumberOfTasks,
	}
}