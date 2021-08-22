package usecase

import (
	"context"
	"errors"
	"log"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/internal/common"
	"github.com/manabie-com/togo/internal/dto"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
)

const (
	jwtPattern      = "^Bearer (\\S*)"
	jwtClaimsUserID = "user_id"
	jwtClaimsExp    = "exp"
)

type (
	UserUsecase interface {
		Login(ctx context.Context, req *dto.LoginRequestDTO) (*dto.LoginResponseDTO, error)
		VerifyToken(ctx context.Context, req *dto.VerifyTokenRequestDTO) (*dto.VerifyTokenResponseDTO, error)
	}

	userUsecase struct {
		JWTKey string
		store  *sqllite.LiteDB
	}
)

func NewLoginUsecase(jwtKey string, store *sqllite.LiteDB) UserUsecase {
	return &userUsecase{
		JWTKey: jwtKey,
		store:  store,
	}
}

func (u *userUsecase) createToken(id string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims[jwtClaimsUserID] = id
	atClaims[jwtClaimsExp] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(u.JWTKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *userUsecase) Login(ctx context.Context, req *dto.LoginRequestDTO) (resp *dto.LoginResponseDTO, err error) {
	if req.UserID == "" || req.Password == "" {
		return nil, errors.New(common.ReasonUserIDPasswordEmptyError.Code())
	}

	if !u.store.ValidateUser(ctx, req.UserID, req.Password) {
		return nil, errors.New(common.ReasonUnauthorized.Code())
	}

	token, err := u.createToken(req.UserID)
	if err != nil {
		return nil, errors.New(common.ReasonCreateTokenError.Code())
	}

	return &dto.LoginResponseDTO{
		Token: token,
	}, nil
}

func (u *userUsecase) VerifyToken(ctx context.Context, req *dto.VerifyTokenRequestDTO) (resp *dto.VerifyTokenResponseDTO, err error) {
	pattern, _ := regexp.Compile(jwtPattern)
	if req.Token == "" || !pattern.Match([]byte(req.Token)) {
		log.Println("token not match pattern: ^Bearer (\\S*)")
		return nil, errors.New(common.ReasonInvalidToken.Code())
	}
	tokens := pattern.FindSubmatch([]byte(req.Token))
	if len(tokens) != 2 || string(tokens[1]) == "" {
		log.Println("token is empty")
		return nil, errors.New(common.ReasonInvalidToken.Code())
	}
	token := string(tokens[1])

	claims := make(jwt.MapClaims)
	jwtToken, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(u.JWTKey), nil
	})
	if err != nil {
		log.Printf("parse with claims error: %v", err)
		return nil, errors.New(common.ReasonInvalidToken.Code())
	}

	if !jwtToken.Valid {
		log.Println("jwtToken invalid")
		return nil, errors.New(common.ReasonInvalidToken.Code())
	}

	id, ok := claims[jwtClaimsUserID].(string)
	if !ok {
		log.Println("claims user id not ok")
		return nil, errors.New(common.ReasonInvalidToken.Code())
	}
	return &dto.VerifyTokenResponseDTO{
		UserID: id,
	}, nil
}
