package models

import "manabie.com/internal/common"

type Task struct {
	Id int
	Title string
	Content string
	CreatedTime common.Time
	Owner *User
}

func MakeTask(
	iId int,
	iTitle string,
	Content string,
	iCreatedTime common.Time,
	iOwner *User,
) Task {
	return Task{
		Id: iId,
		Title: iTitle,
		Content: Content,
		Owner: iOwner,
		CreatedTime: iCreatedTime,
	}
}