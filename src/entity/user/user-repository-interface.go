package user

type IUserRepository interface {
	Create(user *User) (*User, error)
	FindOne(options interface{}) (*User, error)
	Find(options interface{}) (*[]User, error)
	UpdateById(id string, user *User) (*User, error)
	DeleteById(id string) (bool, error)
}
