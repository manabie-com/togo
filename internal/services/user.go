package services

import (
	"context"
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	usermodel "github.com/manabie-com/togo/internal/storages/user/model"
	usersqlstore "github.com/manabie-com/togo/internal/storages/user/sqlstore"
	"github.com/manabie-com/togo/pkg/common/crypto"
	"github.com/manabie-com/togo/pkg/common/xerrors"
	"github.com/manabie-com/togo/up"
	"time"
)

var _ up.UserService = &UserService{}

type UserService struct {
	userstore *usersqlstore.UserStore
	maxTodo   int
	jwtKey    string
}

func NewUserService(db *sql.DB, maxTodo int, jwtKey string) *UserService {
	return &UserService{
		userstore: usersqlstore.NewUserStore(db),
		maxTodo:   maxTodo,
		jwtKey:    jwtKey,
	}
}

func (s *UserService) Register(ctx context.Context, req *up.RegisterRequest) (*up.RegisterResponse, error) {
	if req.MaxTodo == 0 {
		req.MaxTodo = s.maxTodo
	}

	userID := sql.NullString{
		String: req.ID,
		Valid:  true,
	}
	user, err := s.userstore.FindByID(ctx, userID)
	if err != nil && err != sql.ErrNoRows {
		return nil, xerrors.Error(xerrors.Internal, err)
	}

	if user != nil {
		return nil, xerrors.ErrorM(xerrors.InvalidArgument, nil, "user exists with the given id")
	}

	err = s.userstore.Create(ctx, &usermodel.User{
		ID:       req.ID,
		Password: req.Password,
		MaxTodo:  req.MaxTodo,
	})
	if err != nil {
		return nil, xerrors.Error(xerrors.Internal, nil)
	}

	return &up.RegisterResponse{
		ID:      req.ID,
		MaxTodo: req.MaxTodo,
	}, nil
}

func (s *UserService) Login(ctx context.Context, req *up.LoginRequest) (*up.LoginResponse, error) {
	userID := sql.NullString{
		String: req.UserID,
		Valid:  true,
	}

	user, err := s.userstore.FindByID(ctx, userID)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, xerrors.Error(xerrors.Internal, err)
		}
		return nil, xerrors.ErrorM(xerrors.UnAuthorized, err, "incorrect user_id/pwd")
	}

	if user == nil || !crypto.CheckPasswordHash(req.Password, user.Password) {
		return nil, xerrors.ErrorM(xerrors.UnAuthorized, nil, "incorrect user_id/pwd")
	}

	token, err := s.createToken(req.UserID)
	if err != nil {
		return nil, xerrors.Error(xerrors.Internal, err)
	}

	resp := up.LoginResponse(token)
	return &resp, nil
}

func (s *UserService) createToken(id string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(s.jwtKey))
	if err != nil {
		return "", err
	}
	return token, nil
}
