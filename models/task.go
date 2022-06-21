package models

import "time"

type Task struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UserId    int32     `json:"userId"`
}

type Status string

const (
	toDo  Status = "todo"
	doing Status = "doing"
	done  Status = "done"
)

func (e Status) isValidStatus() bool {
	switch e {
	case toDo, doing, done:
		return true
	}
	return false
}

type task interface {
	/* map[string]interface{} like
	{
		status: "success"||"failure",
		message: Optional
		data: {
			data
		}
	}
	*/
	Create() map[string]interface{}
	GetTasks(userId int32) map[string]interface{}
	GetTask(taskId uint32) map[string]interface{}
}
