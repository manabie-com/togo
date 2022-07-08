package business

import (
	"context"
	"errors"
	"fmt"
	"time"
	"todo/be/db"
	"todo/be/utils"

	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson"
)

func Todo_add(ctx context.Context, user db.User, text string) error {
	if len(user.UserId) == 0 {
		return errors.New("Invald UserId")
	}
	if len(text) == 0 {
		return errors.New("Invald Text")
	}
	count := todo_countAddedToday(ctx, user.UserId)
	if count < 0 {
		return errors.New("Unknow Error")
	}
	if count >= user.Limit {
		return errors.New("Limit Task")
	}
	todo := db.Todo{
		UserId:      user.UserId,
		Text:        text,
		CreatedDate: time.Now().Unix(),
	}
	err := db.Todo_add(ctx, todo)
	if !utils.IsError(err) {
		fmt.Println(todo.UserId + " Add todo task success: " + cast.ToString(count+1) + "/" + cast.ToString(user.Limit))
	}
	return err
}

func todo_countAddedToday(ctx context.Context, userId string) int64 {
	return db.Todo_count(ctx, bson.M{
		db.Todo_UserId:      userId,
		db.Todo_CreatedDate: todo_getToDayQuery(time.Now().Unix()),
	})
}

func todo_getToDayQuery(curTime int64) bson.M {
	const oneDaySecs = 86400
	firstSecOfDay := curTime - curTime%oneDaySecs
	return bson.M{
		"$gte": firstSecOfDay,
		"$lt":  firstSecOfDay + oneDaySecs,
	}
}
