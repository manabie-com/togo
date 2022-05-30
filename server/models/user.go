package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email, omitempty"`
	Password string             `json:"password,omitempty" bson:"password, omitempty"`
	Token    string             `json:"token,omitempty" bson:"token, omitempty"`
	Limit    int                `json:"limit,omitempty" bson:"limit, omitempty"`
	Tasks    []Task             `json:"limit,omitempty" bson:"limit, omitempty"`
}
