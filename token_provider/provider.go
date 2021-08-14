package tokenprovider

type Provider interface {
	Generate(id int, expiry int) (*Token, error)
	GenRefreshToken(pl IPayload, expiry int) (*Token, error)
	GenAccessToken(pl IPayload, expiry int) (*Token, error)
	Validate(token string) (IPayload, error)
}
