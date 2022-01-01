package user

import (
	"context"

	"github.com/manabie-com/togo/app/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetOneByUsername get one by username
func (r *repoManager) GetOneByUsername(ctx context.Context, username string) (result model.User, err error) {

	// set condition
	condition := bson.M{}
	condition["username"] = username

	err = r.db.FindOne(ctx, condition).Decode(&result)
	if err == mongo.ErrNoDocuments {
		err = nil
	}

	return
}
