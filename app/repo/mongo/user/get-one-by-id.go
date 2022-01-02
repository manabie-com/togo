package user

import (
	"context"

	"github.com/manabie-com/togo/app/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetOneByID get one by id
func (r *repoManager) GetOneByID(ctx context.Context, id int) (result model.User, err error) {

	// set condition
	condition := bson.M{}
	condition["_id"] = id

	err = r.db.FindOne(ctx, condition).Decode(&result)
	if err == mongo.ErrNoDocuments {
		err = nil
	}

	return
}
