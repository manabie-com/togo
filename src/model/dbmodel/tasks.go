package dbmodel

import "time"

type Status string

const (
	COMPLETE   = "complete"
	INPROGRESS = "inprogress"
	NOTSATRTED = "notstarted"
)

// TaskHistory structure in db
type TaskHistory struct {
	Time    *time.Time `bson:"time" json:"time"`
	Details string     `bson:"details" json:"details"`
}

// Task structure in db
type Task struct {
	Id          string        `bson:"id" json:"id"`
	Title       string        `bson:"title" json:"title"`
	Details     string        `bson:"details" json:"details"`
	Status      Status        `bson:"status" json:"status"`
	TaskHistory []TaskHistory `bson:"taskHistory" json:"taskHistory"`
}
