package services

import (
	"github.com/google/uuid"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	models "github.com/manabie-com/togo/internal/models"
	repositories "github.com/manabie-com/togo/internal/repositories"
)

type AuthService struct {
	UserRepository repositories.UserRepository
	JWTKey string
}

func ProvideAuthService(repo repositories.UserRepository) AuthService {
	return AuthService{UserRepository: repo}
}

func (service *AuthService) Signup(user models.User) (models.User, error) {
	// Implement UUID for user_id
	user.ID = uuid.Must(uuid.NewUUID())
	service.UserRepository.Create(user)

	token, err := service.createToken(user.ID.String())
	user.AccessToken = token
	return user, err
}

func (service *AuthService) Login(user models.User) (models.User, error) {
	// Find user by email & password
	var condition = map[string]interface{}{
		"email": user.Email,
		"password": user.Password,
	}
	user, err := service.UserRepository.FindWhere(condition)

	// create access token
	token, err := service.createToken(user.ID.String())
	user.AccessToken = token

	return user, err
}

func (s *AuthService) createToken(userId string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(s.JWTKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *AuthService) ValidToken(token string) (string, bool) {
	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(s.JWTKey), nil
	})
	if err != nil {
		log.Println(err)
		return "", false
	}

	if !t.Valid {
		return "", false
	}

	userId, ok := claims["user_id"].(string)
	if !ok {
		return "",false
	}

	return userId, true
}
