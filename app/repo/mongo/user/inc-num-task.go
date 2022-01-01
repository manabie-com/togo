package user

import (
	"context"
	"time"

	"github.com/manabie-com/togo/app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// IncNumTaskReq req inc num current task
type IncNumTaskReq struct {
	UserID   int
	MaxTasks int
	// tracing
	UpdatedIP string
}

// IncNumTask increase task of user, must check reach limit or not ?
func (r *repoManager) IncNumTask(ctx context.Context, req IncNumTaskReq) (result model.User, err error) {
	// set condition
	condition := bson.M{}
	condition["_id"] = req.UserID
	condition["current_tasks"] = bson.M{
		"$lt": req.MaxTasks,
	}

	// update
	set := bson.M{}
	set["updated_date"] = time.Now()
	if req.UpdatedIP != "" {
		set["updated_ip"] = req.UpdatedIP
	}
	inc := bson.M{
		"current_tasks": 1,
	}

	// update
	updates := bson.M{
		"$set": set,
		"$inc": inc,
	}

	// set options
	option := options.FindOneAndUpdate()
	option.SetReturnDocument(options.After)

	myResult := r.db.FindOneAndUpdate(ctx,
		condition,
		updates,
		option)

	if myResult.Err() != nil {
		err = myResult.Err()
		if err == mongo.ErrNoDocuments {
			err = nil
		}
		return
	}

	err = myResult.Decode(&result)
	return
}
