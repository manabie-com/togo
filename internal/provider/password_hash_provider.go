package provider

type PasswordHashProvider interface {
	HashPassword(password string) (string, error)
	ComparePassword(password, hashPassword string) error
}
