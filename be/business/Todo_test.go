package business

import (
	"context"
	"testing"
	"time"
	"todo/be/db"
	"todo/be/utils"

	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson"
)

func TestTodo_add(t *testing.T) {
	successDb := db.InitDb()
	if !successDb {
		t.Errorf("Output expect true instead of false")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	user := db.User{
		UserId: "",
		Limit:  0,
	}

	result := Todo_add(ctx, user, "")
	if !utils.IsError(result) {
		t.Errorf("Output expect error instead of nil")
	}
	if result.Error() != "Invald UserId" {
		t.Errorf("Error message expect 'Invald UserId' instead of '%v'", result.Error())
	}

	user.UserId = "UserId"
	result = Todo_add(ctx, user, "")
	if !utils.IsError(result) {
		t.Errorf("Output expect error instead of nil")
	}
	if result.Error() != "Invald Text" {
		t.Errorf("Error message expect 'Invald Text' instead of '%v'", result.Error())
	}

	user.UserId = "UserIdTestAdd"
	db.Todo_delete(ctx, bson.M{db.Todo_UserId: user.UserId})

	result = Todo_add(ctx, user, "Text")
	if !utils.IsError(result) {
		t.Errorf("Output expect error instead of nil")
	}
	if result.Error() != "Limit Task" {
		t.Errorf("Error message expect 'Limit Task' instead of '%v'", result.Error())
	}

	user.Limit = 1
	result = Todo_add(ctx, user, "Text")
	if utils.IsError(result) {
		t.Errorf("Output expect nil instead of error")
	} else {
		result = Todo_add(ctx, user, "Text")
		if !utils.IsError(result) {
			t.Errorf("Output expect error instead of nil")
		}
		if result.Error() != "Limit Task" {
			t.Errorf("Error message expect 'Limit Task' instead of '%v'", result.Error())
		}
	}
	db.Todo_delete(ctx, bson.M{db.Todo_UserId: user.UserId})
}

func TestTodo_countAddedToday(t *testing.T) {
	successDb := db.InitDb()
	if !successDb {
		t.Errorf("Output expect true instead of false")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	user := db.User{
		UserId: "UserIdTestCount",
		Limit:  2,
	}
	db.Todo_delete(ctx, bson.M{db.Todo_UserId: user.UserId})
	count := todo_countAddedToday(ctx, user.UserId)
	if count != 0 {
		t.Errorf("Output expect 0 instead of %d", count)
	}
	Todo_add(ctx, user, "Text")
	count = todo_countAddedToday(ctx, user.UserId)
	if count != 1 {
		t.Errorf("Output expect 1 instead of %d", count)
	}
	Todo_add(ctx, user, "Text")
	count = todo_countAddedToday(ctx, user.UserId)
	if count != 2 {
		t.Errorf("Output expect 2 instead of %d", count)
	}
	Todo_add(ctx, user, "Text")
	count = todo_countAddedToday(ctx, user.UserId)
	if count != 2 {
		t.Errorf("Output expect 2 instead of %d", count)
	}
	db.Todo_delete(ctx, bson.M{db.Todo_UserId: user.UserId})
}

func TestTodo_getToDayQuery(t *testing.T) {
	const curTimeInDay int64 = 1657168378
	const firstSecInDay int64 = 1657152000
	const firstSecNextDay int64 = 1657238400
	result := todo_getToDayQuery(curTimeInDay)

	gteData, okGte := result["$gte"]
	if !okGte {
		t.Errorf("Output expect to have key $gte")
	}
	gte := cast.ToInt64(gteData)
	if gte != firstSecInDay {
		t.Errorf("Output expect %d instead of %d", firstSecInDay, gte)
	}

	ltData, okLt := result["$lt"]
	if !okLt {
		t.Errorf("Output expect to have key $gte")
	}
	lt := cast.ToInt64(ltData)
	if lt != firstSecNextDay {
		t.Errorf("Output expect %d instead of %d", firstSecNextDay, lt)
	}
}
