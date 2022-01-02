package task

import (
	"context"
	"time"

	"github.com/manabie-com/togo/app/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateReq create user request
type CreateReq struct {
	UserID      int
	Name        string
	Description string
	// tracing
	CreatedIP string
}

// Create : create new task
func (r *repoManager) Create(ctx context.Context, req CreateReq) (model.Task, error) {
	// create model insert to db
	now := time.Now()
	taskInsert := model.Task{
		ID:          primitive.NewObjectID(),
		UserID:      req.UserID,
		Name:        req.Name,
		Description: req.Description,
		Status:      model.TaskStatusActive,
		// tracing
		CreatedIP:   req.CreatedIP,
		CreatedDate: &now,
	}

	_, err := r.db.InsertOne(ctx, taskInsert)
	return taskInsert, err
}
