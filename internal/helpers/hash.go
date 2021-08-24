package helpers

import "golang.org/x/crypto/bcrypt"

type PasswordProvider struct {
}

func (p *PasswordProvider) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (p *PasswordProvider) CheckPasswordHash(password, hash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}
	return true
}
