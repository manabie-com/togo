package user

import (
	"context"
	"time"

	"github.com/manabie-com/togo/app/model"
)

// CreateReq create user request
type CreateReq struct {
	Username       string
	HashedPassword string
	MaxTasks       int
	// tracing
	CreatedIP string
}

// Create : create new property
func (r *repoManager) Create(ctx context.Context, req CreateReq) (model.User, error) {
	// create model insert to db
	now := time.Now()
	userInsert := model.User{
		ID:             r.db.Mongo.NextID(r.db.Collection.Name()),
		Username:       req.Username,
		HashedPassword: req.HashedPassword,
		MaxTasks:       req.MaxTasks,
		CurrentTasks:   0,
		Status:         model.UserStatusActive,
		// tracing
		CreatedIP:   req.CreatedIP,
		CreatedDate: &now,
	}

	_, err := r.db.InsertOne(ctx, userInsert)
	return userInsert, err
}
