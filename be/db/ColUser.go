package db

type User struct {
	UserId string `json:"UserId,omitempty" bson:"UserId,omitempty"`
	Limit  int64  `json:"Limit,omitempty" bson:"Limit,omitempty"`
}
