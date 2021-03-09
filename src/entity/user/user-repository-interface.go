package user

type IUserRepository interface {
	FindOne(options interface{}) (*User, error)
}
