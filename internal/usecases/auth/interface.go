package auth

type Common interface {
	ValidateUser(username string) (bool, error)
	GenerateToken(userID, maxTaskPerday string) (*string, error)
}

type AuthRepository interface {
	Common
}
