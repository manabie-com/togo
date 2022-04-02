package hasher

import "golang.org/x/crypto/bcrypt"

type BcryptProvider struct {
	salt string
	cost int
}

func NewBcryptProvider(salt string, cost int) *BcryptProvider {
	return &BcryptProvider{
		salt: salt,
		cost: cost,
	}
}

func (b *BcryptProvider) HashPassword(password string) (string, error) {
	arr, err := bcrypt.GenerateFromPassword([]byte(password), b.cost)
	if err != nil {
		return "", err
	}
	return string(arr), nil
}

func (b *BcryptProvider) ComparePassword(password, hashPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}