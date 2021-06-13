package services

type AuthService interface {
	Login(username string, pwd string) (string, error)
}

type authServiceImpl struct {
	JWTKey string
}

func (a *authServiceImpl) Login(username string, pwd string) (string, error) {
	return "", nil
}
