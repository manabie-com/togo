package user

type IUserRepository interface {
	FindFirstByID(id int64) (User, error)
	FindFirstByUsername(username string) (User, error)
}
