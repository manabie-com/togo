package auth

import "github.com/ansidev/togo/domain/user"

type ICredRepository interface {
	Save(userModel user.User) (string, error)
	Get(token string) (AuthenticationCredential, error)
}
