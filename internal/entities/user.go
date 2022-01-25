package entities

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	TaskLimit int    `json:"taskLimit"`
	Jwt       string `json:"jwt"`
}
