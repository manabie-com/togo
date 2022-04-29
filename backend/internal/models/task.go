package models

import "manabie.com/internal/common"

type Task struct {
	Id int 						`json:"id"`
	Title string 				`json:"title"`
	Content string  			`json:"content"`
	CreatedTime common.Time  	`json:"created_time"`
	Owner *User					`json:"-"`
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