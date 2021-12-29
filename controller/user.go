package controller

import (
	"togo/db"
	"togo/model"
)

func CreateUser(input *model.User)  {
	db.DB.Create(input)
}
