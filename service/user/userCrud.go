package user

import (
	"crypto/sha256"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"manabie.com/togo"
	"manabie.com/togo/utils"
)

type Config struct {
	Store      togo.Store
	SigningKey string
}

type userCrudServiceImpl struct {
	store      togo.Store
	signingKey string
}

func New(cfg Config) togo.UserCrudService {
	return &userCrudServiceImpl{
		store:      cfg.Store,
		signingKey: cfg.SigningKey,
	}
}

func (u userCrudServiceImpl) CreateUser(rq togo.UserRequest) (togo.UserResponse, error) {
	entity, err := u.store.CreateUser(rq.Username, hashPassword(rq.Password))
	if err != nil {
		return togo.UserResponse{}, errors.Wrapf(err, "fail to create user")
	}

	return togo.UserResponse{Id: entity.Id, Username: entity.Username, CreatedAt: entity.CreatedAt, LastLogin: entity.LastLogin}, nil
}

func (u userCrudServiceImpl) Login(rq togo.UserRequest) (string, error) {
	userId, err := u.store.Login(rq.Username, hashPassword(rq.Password))
	if err != nil {
		return "", errors.Wrapf(err, "error when login user")
	}

	now := utils.NowInUnixSecond()
	claims := customClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now,
			ExpiresAt: now + 3600,
			Issuer:    "manabie",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(u.signingKey))
	if err != nil {
		return "", errors.Wrapf(err, "error when create jwt for user")
	}
	return signedToken, nil
}

func hashPassword(pwd string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(pwd)))
}

type customClaims struct {
	UserId int `json:"userId"`
	jwt.StandardClaims
}
