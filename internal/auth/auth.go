package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jmsemira/togo/internal/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var SALT string = "This is a sample super secret salt for password encryption"
var bcryptDiff = 12

// HashPass Generates a hash from a password and salt
func HashPass(password string) string {
	saltedPassword := []byte(password + SALT)
	hash, err := bcrypt.GenerateFromPassword(saltedPassword, bcryptDiff)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	return string(hash)
}

func Login(username, password string) (*models.User, error) {
	saltedPassword := []byte(password + SALT)
	user := models.User{}
	user.GetUserbyUsername(username)

	hashedPassword := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(hashedPassword, saltedPassword)

	if err != nil || user.ID == 0 {
		return nil, errors.New("Invalid username and password")
	}
	return &user, nil
}

func Register(user *models.User) error {
	if user.UsernameExist() {
		return errors.New("Username Exist!")
	}

	user.Password = HashPass(user.Password)
	user.Save()
	return nil
}

func GenerateJWTToken(user *models.User) (string, error) {
	expirationTime := GenerateExpirationTime()

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		ID:        user.ID,
		Username:  user.Username,
		RateLimit: user.RateLimitPerDay,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	})

	token, err := claims.SignedString([]byte("this is my secret key"))

	if err != nil {
		return "", err
	}

	return token, nil
}

var GenerateExpirationTime = func() time.Time {
	return time.Now().Add(7 * 24 * 60 * time.Minute)
}
