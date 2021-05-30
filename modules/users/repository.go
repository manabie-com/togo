package users

type UsersRepository interface {
	CheckLogin(userId string, pass string) (Users, error)
}
