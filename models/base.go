package models

const (
	DEFAULT_MAX_DAILY_TASKS uint = 16
)

// Models structure contains all models informations
type Models struct {
	User *UserModel
	Task *TaskModel
}

// UserTask DTO
type UserTask struct {
	UserID        string `json:"userId" required:"true"`
	MaxDailyLimit uint   `json:"maxDailyLimit"`
	TodoTask      string `json:"task" required:"true"`
}
