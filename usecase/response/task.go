package response

import "github.com/go-kit/kit/endpoint"

var (
	_ endpoint.Failer = &GetTasks{}
	_ endpoint.Failer = &CreateTask{}
)

type TaskData struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	UserID      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}

type GetTasks struct {
	Data []*TaskData `json:"data"`
	Err  error       `json:"-"`
}

func (r GetTasks) Failed() error {
	return r.Err
}

type CreateTask struct {
	Data *TaskData `json:"data"`
	Err  error     `json:"-"`
}

func (r CreateTask) Failed() error {
	return r.Err
}
