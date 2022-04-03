package crypter

import "golang.org/x/crypto/bcrypt"

// Crypter holds crypter methods
type Crypter struct{}

// New creates crypter service
func New() *Crypter {
	return &Crypter{}
}

func (*Crypter) HashPassword(password string) string {
	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPwd)
}

func (*Crypter) CompareHashAndPassword(hashedPwd, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(password)) == nil
}
