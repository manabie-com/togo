package services

type UsersService interface {
	ValidateUser(userID, pwd string) bool
	Login(userID, pwd string) (token string, err error, code int)
}
