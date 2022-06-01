package models

type User struct {
	ID       string `json:"id,omitempty" bson:"_id,omitempty"`
	Email    string `json:"email,omitempty" bson:"email, omitempty"`
	Password string `json:"password,omitempty" bson:"password, omitempty"`
	Token    string `json:"token,omitempty" bson:"token, omitempty"`
	Limit    int    `json:"limit,omitempty" bson:"limit, omitempty"`
}
