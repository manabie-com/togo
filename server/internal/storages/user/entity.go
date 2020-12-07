package user

// User reflects users data from DB
type User struct {
	ID       uint64
	Name string
	Password string
}
