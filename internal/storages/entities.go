package storages

// Task reflects tasks in DB
type Task struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	UserID      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}

// User reflects users data from DB
type User struct {
	ID       string
	Password string
}

// LoginJSONResponse represents the response from a HTTP GET to /login
type LoginJSONResponse struct {
	Data       string `json:"data"`
	Error      string `json:"error"`
	StatusCode int    `json:"-"`
}

// ListTasksJSONResponse represents the response from a HTTP GET to /tasks
type ListTasksJSONResponse struct {
	Data       []*Task `json:"data"`
	Error      string  `json:"error"`
	StatusCode int     `json:"-"`
}

// AddTasksJSONResponse represents the response from a HTTP POST to /tasks
type AddTasksJSONResponse struct {
	Data       *Task  `json:"data"`
	Error      string `json:"error"`
	StatusCode int    `json:"-"`
}
