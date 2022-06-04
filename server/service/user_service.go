package service

import (
	"errors"
	"math/rand"
	"os"
	"strconv"
	"time"
	"togo/common/key"
	"togo/models"
	"togo/utils/security"

	repository "togo/repository/user"

	"github.com/golang-jwt/jwt"
	"github.com/rs/xid"
)

// Define `User` Service Interface with the following
// Methods which will be utilized by the `UserController`
type UserService interface {
	// Generate ID, generate random limit, hash password
	Register(user *models.User) (*models.User, error)

	// Pass the `user` to the User repository
	Login(user *models.User) error

	// Check for missing fields upon registration
	ValidateRegistration(user *models.User) error

	// Check for missing fields upon login
	ValidateLogin(user *models.User) error

	// Set claims and generate JWT token
	GenerateJWT(user *models.User, expiration time.Time) (string, error)
}

// Define Struct with UserRepository as the attribute
// This attribute is responsible for `User` database interactions
type userservice struct {
	userrepository repository.UserRepository
}

// Define a Constructor
// Dependency Injection for `User` Service
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

// Validate Login `User`
func (s *userservice) ValidateLogin(user *models.User) error {
	// Check if user exists in the database
	foundUser, err := s.userrepository.GetUserByEmail(user.Email, user)
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

// Register a new `User`
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
	tokenString, err := token.SignedString(key.JwtKey)
	if err != nil {
		return "", errors.New("unable to generate JWT token")
	}

	// Token
	user.Token = tokenString
	return tokenString, nil
}

// Login existing `User`
func (s *userservice) Login(user *models.User) error {
	return s.userrepository.Login(user)
}
