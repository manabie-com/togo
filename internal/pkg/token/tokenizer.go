package token

type Tokenizer interface {
	Create(payload *Payload) (string, error)
	Verify(token string) (*Payload, error)
}
