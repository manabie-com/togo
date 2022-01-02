package task

import (
	"context"

	"github.com/manabie-com/togo/app/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AllReq get all flow request
type AllReq struct {
	UserID *int
	Status *model.TaskStatus
	Offset primitive.ObjectID
	Limit  int
}

// All get all property
func (r *repoManager) All(ctx context.Context, req AllReq) (results []model.Task, err error) {

	// set options
	option := options.Find()
	option.SetSort(bson.M{"_id": -1})

	if req.Limit != -1 {
		// set default
		option.SetLimit(100)
		if req.Limit > 0 && req.Limit <= 100 {
			option.SetLimit(int64(req.Limit))
		}
	}

	// set filter
	filter := bson.M{}
	if req.UserID != nil {
		filter["user_id"] = *req.UserID
	}
	if req.Status != nil {
		filter["status"] = *req.Status
	}
	if !req.Offset.IsZero() {
		filter["_id"] = bson.M{"$lt": req.Offset}
	}

	cur, err := r.db.Find(ctx, filter, option)
	if err != nil {
		return
	}
	// Close the cursor once finished
	defer cur.Close(ctx)

	err = cur.All(ctx, &results)
	return
}
