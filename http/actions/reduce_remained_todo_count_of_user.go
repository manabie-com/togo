package actions

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"pt.example/grcp-test/database"
	"pt.example/grcp-test/models"
)

type ReduceRemainedTodoCountOfUserParam interface {
	GetAssigneeEmail() string // Used to get user
	GetTaskSavedCount() uint8 // Used to reduce remained todo task per day
}

func ReduceRemainedTodoCountOfUser(ctx context.Context, p ReduceRemainedTodoCountOfUserParam) (r *mongo.UpdateResult, err error) {
	var u models.User

	var ur database.Repository = &u

	today := time.Now()
	f := ur.GetCollection().FindOneAndUpdate(ctx, bson.D{{Key: "email", Value: p.GetAssigneeEmail()}}, bson.D{{Key: "$set", Value: bson.D{{Key: "in_progress", Value: true}}}})

	if f.Err() != nil {
		err = f.Err()
		return
	}

	f.Decode(&u)

	if u.InProgress {
		err = errors.New("The user in another process")
		return
	}

	if isNotSameDay(today.Local(), u.LastAssignedTime.Time().Local()) {
		u.RemainedAssignableTaskPerDay = u.MaxAssignedTaskPerDay
	}

	if u.RemainedAssignableTaskPerDay == 0 {
		err = errors.New("Tasks of this user was full")
		return
	}

	u.RemainedAssignableTaskPerDay -= p.GetTaskSavedCount()
	u.LastAssignedTime = primitive.NewDateTimeFromTime(today)

	r, err = ur.GetCollection().UpdateByID(ctx, u.Id, bson.D{{Key: "$set", Value: u}})

	return
}

func isNotSameDay(dt1 time.Time, dt2 time.Time) bool {
	y1, m1, d1 := dt1.Date()
	y2, m2, d2 := dt2.Date()

	if y1 != y2 {
		return true
	}

	if m1 != m2 {
		return true
	}

	if d1 != d2 {
		return true
	}

	return false
}
