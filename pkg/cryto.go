package pkg

import "golang.org/x/crypto/bcrypt"

type appCrypto struct {
}

func NewCrypto() *appCrypto {
	return &appCrypto{}
}

func (a *appCrypto) HashPassword(password string) string {
	return hashPassword(password)
}

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func (a appCrypto) ComparePassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
