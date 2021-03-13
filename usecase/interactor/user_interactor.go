package interactor

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"

	"github.com/valonekowd/togo/domain/entity"
	"github.com/valonekowd/togo/usecase/interfaces"
	"github.com/valonekowd/togo/usecase/request"
	"github.com/valonekowd/togo/usecase/response"
)

var (
	ErrEmailAlreadyExists     = errors.New("email already exists")
	ErrEmailNotFound          = errors.New("email not found")
	ErrInvalidEmailOrPassword = errors.New("invalid email or password")
)

type UserInteractor interface {
	SignUp(context.Context, request.SignUp) (*response.SignUp, error)
	SignIn(context.Context, request.SignIn) (*response.SignIn, error)
}

func NewUserInteractor(ds interfaces.DataSource, presenter interfaces.UserPresenter, logger log.Logger) UserInteractor {
	var i UserInteractor
	{
		i = NewBasicUserInteractor(ds, presenter)
		// i = LoggingInterceptor(logger)(i)
	}
	return i
}

type basicUserInteractor struct {
	ds        interfaces.DataSource
	presenter interfaces.UserPresenter
}

var _ UserInteractor = basicUserInteractor{}

func NewBasicUserInteractor(ds interfaces.DataSource, presenter interfaces.UserPresenter) UserInteractor {
	return basicUserInteractor{
		ds:        ds,
		presenter: presenter,
	}
}

func (i basicUserInteractor) SignUp(ctx context.Context, req request.SignUp) (*response.SignUp, error) {
	u := &entity.User{
		ID:        uuid.New().String(),
		Email:     req.Email,
		Password:  req.Password,
		MaxTodo:   req.MaxTodo,
		CreatedAt: time.Now().UTC(),
	}

	if err := u.HashPassword(); err != nil {
		return nil, fmt.Errorf("hashing user password: %w", err)
	}

	if err := i.ds.User.Add(ctx, u); err != nil {
		if err, ok := err.(*pq.Error); ok && err.Code.Name() == "unique_violation" && err.Constraint == "users_email_key" {
			return nil, fmt.Errorf("adding user: %w", ErrEmailAlreadyExists)
		}
		return nil, fmt.Errorf("adding user: %w", err)
	}

	return i.presenter.SignUp(ctx, u)
}

func (i basicUserInteractor) SignIn(ctx context.Context, req request.SignIn) (*response.SignIn, error) {
	u, err := i.ds.User.FindByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("finding user: %w", ErrEmailNotFound)
		}
		return nil, fmt.Errorf("finding user: %w", err)
	}

	if isMatch, err := u.ComparePassword(req.Password); !isMatch {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, fmt.Errorf("comparing user password: %w", ErrInvalidEmailOrPassword)
		}
		return nil, fmt.Errorf("comparing user password: %w", err)
	}

	return i.presenter.SignIn(ctx, u)
}
