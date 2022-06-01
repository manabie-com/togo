package service

import (
	"errors"
	"math/rand"
	"os"
	"strconv"
	"time"
	"togo/models"
	"togo/service/key"
	"togo/utils/security"

	repository "togo/repository/user"

	"github.com/golang-jwt/jwt"
	"github.com/rs/xid"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

// Define User Service Interface
type UserService interface {
	Register(user *models.User) (*models.User, error)
	Login(user *models.User) error
	ValidateRegistration(user *models.User) error
	ValidateLogin(user *models.User) error
	GenerateJWT(user *models.User, expiration time.Time) (string, error)
}

// Define Struct with UserRepository as the attribute
type userservice struct {
	userrepository repository.UserRepository
}

// Define a Constructor
// Dependency Injection for User Service
func NewUserService(repo repository.UserRepository) UserService {
	return &userservice{
		userrepository: repo,
	}
}

// Validate the User registration
func (s *userservice) ValidateRegistration(user *models.User) error {
	if user == nil {
		return errors.New("the user is empty")
	}
	if user.Email == "" {
		return errors.New("email is empty")
	}
	if user.Password == "" {
		return errors.New("password is empty")
	}
	return nil
}

// Validate Login user
func (s *userservice) ValidateLogin(user *models.User) error {
	// Check if user exists in the database
	foundUser, err := s.userrepository.GetUser(user.Email, user)
	if err != nil {
		return errors.New("user not found")
	}

	// Check if password provided by the user matches the password in the database
	if ok := security.CheckPasswordHash(user.Password, foundUser.Password); !ok {
		return errors.New("password incorrect")
	}

	user.ID = foundUser.ID
	return nil
}

// Register new users
func (s *userservice) Register(user *models.User) (*models.User, error) {
	// Generate xid using 3rd party library
	guid := xid.New()

	// Get only id string
	user.ID = guid.String()

	// Generate random number to indicate daily limit
	limit, _ := strconv.Atoi(os.Getenv("LIMIT"))
	user.Limit = rand.Intn(limit)

	// Hash password using bcrypt
	password, err := security.HashPassword(user.Password)
	if err != nil {
		return user, errors.New("unable to hash password")
	}
	user.Password = password

	return s.userrepository.Register(user)
}

// Generate JWT token
func (s *userservice) GenerateJWT(user *models.User, expiration time.Time) (string, error) {

	// Generate claims
	claims := &key.Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
		},
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", errors.New("unable to generate JWT token")
	}

	// Token
	user.Token = tokenString
	return tokenString, nil
}

// Login existing users
func (s *userservice) Login(user *models.User) error {
	return s.userrepository.Login(user)
}
