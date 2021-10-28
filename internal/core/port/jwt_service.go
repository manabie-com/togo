package port

type JwtService interface {
	CreateToken(userId string) (string, error)
	ParseToken(token string) (string, error)
}
