package tokenprovider

type Provider interface {
	GenAccessToken(pl IPayload, expiry int) (*Token, error)
	Validate(token string) (IPayload, error)
}
