package requestmodel

type CreateTaskRequest struct {
	Title   string `bson:"title" json:"title"`
	Details string `bson:"details" json:"details"`
	Token   string `bson:"token" json:"token"`
}

type UpdateTaskRequest struct {
	Id      string `bson:"id" json:"id"`
	Status  string `bson:"status,omitempty" json:"status,omitempty"`
	Details string `bson:"details" json:"details"`
	Token   string `bson:"token" json:"token"`
}
