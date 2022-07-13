package main

import (
	"context"

	"github.com/joho/godotenv"
	"pt.example/grcp-test/database"
	"pt.example/grcp-test/models"
)

func main() {
	var err error

	err = godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	err = database.Init()
	if err != nil {
		panic(err)
	}

	createUser()
	// dropTasks()
}

func createUser() {
	var ur database.Repository

	us := []models.User{
		{
			Email:                        "ptrung@manabie.test",
			MaxAssignedTaskPerDay:        5,
			RemainedAssignableTaskPerDay: 5,
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

func dropTasks() {
	t := &models.Task{}
	var tr database.Repository = t

	tr.GetCollection().Drop(context.TODO())
}
