package usecase

import (
	"context"
	"errors"
	"log"
	"regexp"

	"github.com/manabie-com/togo/internal/adapter"
	"github.com/manabie-com/togo/internal/common"
	"github.com/manabie-com/togo/internal/dto"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	"github.com/manabie-com/togo/internal/util"
)

const (
	jwtPattern = "^Bearer (\\S*)"
)

type (
	UserUsecase interface {
		Login(ctx context.Context, req *dto.LoginRequestDTO) (*dto.LoginResponseDTO, error)
		VerifyToken(ctx context.Context, req *dto.VerifyTokenRequestDTO) (*dto.VerifyTokenResponseDTO, error)
	}

	userUsecase struct {
		jwtAdapter adapter.JWTAdapter
		liteDB     sqllite.LiteDB
	}
)

func NewUserUsecase(jwtAdapter adapter.JWTAdapter, liteDB sqllite.LiteDB) UserUsecase {
	return &userUsecase{
		jwtAdapter: jwtAdapter,
		liteDB:     liteDB,
	}
}

func (u *userUsecase) Login(ctx context.Context, req *dto.LoginRequestDTO) (resp *dto.LoginResponseDTO, err error) {
	if req.UserID == "" || req.Password == "" {
		return nil, errors.New(common.ReasonUserIDPasswordEmptyError.Code())
	}

	pwdMd5 := util.GetMD5Hash(req.Password)
	if !u.liteDB.ValidateUser(ctx, req.UserID, pwdMd5) {
		return nil, errors.New(common.ReasonUnauthorized.Code())
	}

	token, err := u.jwtAdapter.CreateToken(ctx, req.UserID)
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

	userID, err := u.jwtAdapter.VerifyToken(ctx, token)
	if err != nil {
		return nil, err
	}

	return &dto.VerifyTokenResponseDTO{
		UserID: userID,
	}, nil
}
