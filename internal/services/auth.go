package services

import (
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/internal/entities"
	"github.com/manabie-com/togo/internal/storages"
	"golang.org/x/crypto/bcrypt"
)

// AuthService defines set of methods or properties needed
// to be implement that will be injected as dependency in
// other components
type AuthService interface {
	DecodeToken(ctx context.Context, token string) (userID string, err error)
	AuthenticateUser(ctx context.Context, userID, password string) (string, error)
}

const (
	// DefaultJWTKey ...
	DefaultJWTKey = "s3cr3t"
	// DefaultPwdHashRound ...
	DefaultPwdHashRound = 10
)

// AuthSvc ...
type AuthSvc struct {
	userRepo     storages.UserRepository
	jwtSecretKey string
	pwdHashRound int
}

// AuthServiceConfiguration ...
type AuthServiceConfiguration struct {
	UserRepo     storages.UserRepository
	JWTKey       string
	PwdHashRound int
}

// NewAuthService ...
func NewAuthService(config AuthServiceConfiguration) *AuthSvc {
	if len(config.JWTKey) == 0 {
		config.JWTKey = DefaultJWTKey
	}
	if config.PwdHashRound == 0 {
		config.PwdHashRound = DefaultPwdHashRound
	}
	return &AuthSvc{
		userRepo:     config.UserRepo,
		jwtSecretKey: config.JWTKey,
		pwdHashRound: DefaultPwdHashRound,
	}
}

// DecodeToken decode the token and return user_id
func (as *AuthSvc) DecodeToken(ctx context.Context, token string) (string, error) {
	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(as.jwtSecretKey), nil
	})
	if err != nil || !t.Valid {
		return "", errors.New("Unauthorized")
	}

	id, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("Unauthorized")
	}
	return id, nil
}

// FindUserByIDNPwd ...
func (as *AuthSvc) FindUserByIDNPwd(ctx context.Context, userID, password string) (*entities.User, error) {
	if len(userID) == 0 || len(password) == 0 {
		return nil, ErrServiceUnhandledException
	}
	foundUser, err := as.userRepo.GetUserByUserID(ctx, userID)
	if foundUser == nil && err != nil {
		return nil, errors.New("Internal error")
	}
	if foundUser == nil && err == nil { // Not found
		return nil, nil
	}
	isPwdMatched := as.ComparePassword(password, foundUser.Password)
	if !isPwdMatched {
		return nil, nil // Password does not match
	}
	return foundUser, nil
}

// AuthenticateUser ...
func (as *AuthSvc) AuthenticateUser(ctx context.Context, userID, password string) (string, error) {
	foundUser, err := as.FindUserByIDNPwd(ctx, userID, password)
	if foundUser == nil {
		if err == nil {
			return "", nil // Not found
		}
		return "", err // Error
	}
	token, err := as.SignJWT(ctx, foundUser.ID)
	if err != nil {
		return "", errors.New("Internal error, cannot sign token")
	}
	return token, nil
}

// SignJWT ...
func (as *AuthSvc) SignJWT(ctx context.Context, userID string) (string, error) {
	if len(userID) == 0 {
		return "", ErrServiceUnhandledException
	}
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userID
	atClaims["exp"] = time.Now().Add(time.Hour * 30).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(as.jwtSecretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

// ComparePassword ...
func (as *AuthSvc) ComparePassword(rawPwd, encryptedPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPwd), []byte(rawPwd))
	return err == nil
}

// hashPassword function to create hashed password before inserting it into database
// func (as *AuthSvc) hashPassword(password string) (string, error) {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
// 	return string(bytes), err
// }
