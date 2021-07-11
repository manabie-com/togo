package token

const (
	UserIDField  = "user_id"
	ExpiredField = "exp"
)

type Generator interface {
	CreateToken(userID string) (string, error)
	ValidateToken(token string) (string, error)
}
