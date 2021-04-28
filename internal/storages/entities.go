package storages

// Task reflects tasks in DB
type Task struct {
	ID          string `pg:"id,type:text,pk" json:"id"`
	Content     string `pg:"content,type:text" json:"content"`
	UserID      string `pg:"user_id,type:text" json:"user_id"`
	CreatedDate string `pg:"created_at,type:text" json:"created_date"`
}

func (Task) TableName() string {
	return "tasks"
}

// User reflects users data from DB
type User struct {
	ID       string `pg:"id,type:text,pk" json:"id"`
	Password string `pg:"password,type:text" json:"password"`
	MaxTodo  int    `pg:"max_todo,type:text" json:"max_todo"`
}

func (User) TableName() string {
	return "users"
}

type EntityUseCase interface {
	GetAuthToken(args map[string]string) (map[string]interface{}, error)
	ListTasks(args map[string]string) (map[string]interface{}, error)
	AddTask(task Task) (map[string]interface{}, error)
}

type EntityRepository interface {
	Login(userID string, pwd string) error
	GetUserByID(userID string) (User, error)
	GetTasks(userID string, createdDate string) ([]Task, error)
	InsertTask(task *Task) error
}
