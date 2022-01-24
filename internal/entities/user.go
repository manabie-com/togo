package entities

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	TaskLimit int64  `json:"taskLimit"`
}
