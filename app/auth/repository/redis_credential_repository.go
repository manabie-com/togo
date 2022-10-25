package repository

import (
	"ansidev.xyz/pkg/rds"
	"encoding/json"
	"github.com/ansidev/togo/domain/auth"
	"github.com/ansidev/togo/domain/user"
	"github.com/ansidev/togo/errs"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
)

func NewRedisCredentialRepository(rdb *rds.RedisDB, tokenTTL time.Duration) auth.ICredRepository {
	return &redisCredentialRepository{db: rdb, tokenTTL: tokenTTL}
}

type redisCredentialRepository struct {
	db       *rds.RedisDB
	tokenTTL time.Duration
}

func (r *redisCredentialRepository) Get(token string) (auth.AuthenticationCredential, error) {
	bytes, err := r.db.GetAsBytes(token)

	if err != nil {
		return auth.AuthenticationCredential{}, errors.Wrap(errs.ErrRecordNotFound, err.Error())
	}

	var authenticationCredential auth.AuthenticationCredential

	if err1 := json.Unmarshal(bytes, &authenticationCredential); err1 != nil {
		return auth.AuthenticationCredential{}, errors.Wrap(errs.ErrInternalAppFailure, err1.Error())
	}

	return authenticationCredential, nil
}

func (r *redisCredentialRepository) Save(userModel user.User) (string, error) {
	authenticationCredential := auth.AuthenticationCredential{
		ID:           userModel.ID,
		MaxDailyTask: userModel.MaxDailyTask,
	}

	bytes, err1 := json.Marshal(authenticationCredential)

	if err1 != nil {
		return "", errors.Wrap(errs.ErrInternalAppFailure, err1.Error())
	}

	token := uuid.NewString()

	_, err2 := r.db.Set(token, bytes, r.tokenTTL)

	if err2 != nil {
		return "", errors.Wrap(errs.ErrDatabaseFailure, err2.Error())
	}

	return token, nil
}
