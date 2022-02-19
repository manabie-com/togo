package provider

// TokenProvider service interface
type TokenProvider interface {
	GenerateToken(data interface{}) (token string, err error)
	VerifyToken(token string) (payload interface{}, err error)
}
