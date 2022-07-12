package actions

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"pt.example/grcp-test/database"
	"pt.example/grcp-test/models"
)

type ReduceRemainedTodoCountOfUserParam interface {
	GetAssigneeEmail() string // Used to get user
	GetTaskSavedCount() uint8 // Used to reduce remained todo task per day
}

func ReduceRemainedTodoCountOfUser(p ReduceRemainedTodoCountOfUserParam) (ok bool) {
	ok = true

	var u models.User

	var ur database.Repository = &u

	r := ur.GetCollection().FindOne(context.TODO(), bson.D{{Key: "email", Value: p.GetAssigneeEmail()}})

	r.Decode(&u)

	u.RemainedAssignableTaskPerDay -= p.GetTaskSavedCount()

	fmt.Printf("%+v\n", u)

	ur.GetCollection().UpdateByID(context.TODO(), u.Id, bson.D{{Key: "$set", Value: u}})

	return
}
