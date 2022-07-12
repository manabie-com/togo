package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"pt.example/grcp-test/database"
	"pt.example/grcp-test/models"
)

func main() {
	database.Init()

	createUser()

	// var u models.User
	// var ur database.Repository = &u
	// r := ur.GetCollection().FindOne(context.TODO(), bson.D{{"email", "ptrung@manabie.test"}})
	// r.Decode(&u)

	// println(u.LastAssignedTime.Time().Local().Date())
}

func createUser() {
	var ur database.Repository

	us := []models.User{
		{
			Email:                 "ptrung@manabie.test",
			MaxAssignedTaskPerDay: 5,
			LastAssignedTime:      primitive.NewDateTimeFromTime(time.Now()),
		},
	}

	ur = &us[0]

	rcs := make([]interface{}, len(us))

	for i, u := range us {
		rcs[i] = u
	}

	ur.GetCollection().Drop(context.TODO())
	ur.GetCollection().InsertMany(context.TODO(), rcs)
}
