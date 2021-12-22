package domain

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shanenoi/togo/common/bcrypt_lib"
	"github.com/shanenoi/togo/common/jwt_lib"
	"github.com/shanenoi/togo/config"
	"github.com/shanenoi/togo/internal/storages/models"
	"github.com/shanenoi/togo/internal/storages/postgresql"
)

type UserDomain interface {
	LoginUser(user *models.User) (token string, err error)
	CheckUserToken(token string) (err error)
	SignupUser(user *models.User) (err error)
}

func NewUserDomain(ctx *gin.Context) UserDomain {
	return &userDomain{ctx}
}

type userDomain struct {
	ctx *gin.Context
}

func(ud *userDomain) LoginUser(user *models.User) (token string, err error) {
	inputedPassword := user.Password

	userPostgreSQL := postgresql.NewUserPostgreSQL(ud.ctx)
	user, err = userPostgreSQL.FindUserByUserName(user.Username)

	correct := bcrypt_lib.CompareHashAndPassword(user.Password, inputedPassword)

	if correct {
		token, err = jwt_lib.Encrypt(models.JwtClaims{user.ID})
	} else {
		err = fmt.Errorf(config.BCRYPT_WRONG_PASSWORD)
	}

	return token, err
}

func(ud *userDomain) CheckUserToken(token string) (err error) {
	claims := map[string]interface{}{}

	claims, err = jwt_lib.Decrypt(token)
	if err != nil {
		return 
	}

	ud.ctx.Set(config.HEADER_USER_ID, claims["user_id"])
	return 
}

func(ud *userDomain) SignupUser(user *models.User) (err error) {
	user.Password, err = bcrypt_lib.GenerateFromPassword(user.Password)
	if err != nil {
		return
	}

	userPostgreSQL := postgresql.NewUserPostgreSQL(ud.ctx)
	err = userPostgreSQL.CreateUser(user)
	return
}
