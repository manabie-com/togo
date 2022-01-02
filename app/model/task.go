package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TaskStatus contain user's status
type TaskStatus string

// TaskStatus define
const (
	TaskStatusActive   TaskStatus = "active"
	TaskStatusInActive TaskStatus = "in_active"
	TaskStatusDelete   TaskStatus = "delete"
)

// Task contain info task of user
type Task struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID      int                `json:"user_id" bson:"user_id"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Status      TaskStatus         `json:"status,omitempty" bson:"status,omitempty"`
	// Tracing
	UpdatedIP   string     `json:"updated_ip,omitempty" bson:"updated_ip,omitempty"`
	UpdatedDate *time.Time `json:"updated_date,omitempty" bson:"updated_date,omitempty"`
	//
	CreatedIP   string     `json:"created_ip,omitempty" bson:"created_ip,omitempty"`
	CreatedDate *time.Time `json:"created_date,omitempty" bson:"created_date,omitempty"`
}

// IsExists check task's existed
func (m *Task) IsExists() bool {
	return !m.ID.IsZero()
}

// IsActive check task's active or not
func (m *Task) IsActive() bool {
	return m.Status == TaskStatusActive
}
