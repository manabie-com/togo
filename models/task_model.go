package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	Id_user   primitive.ObjectID `json:"id_user,omitempty"`
	Name      string             `json:"name,omitempty" validate:"required"`
	CreatedAt string             `json:"created_at,omitempty"`
	Content   string             `json:"content,omitempty" validate:"required"`
	Status    string             `json:"status,omitempty"`
}
