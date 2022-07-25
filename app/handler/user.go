package handler

import (
	"github.com/golang-jwt/jwt"
	"github.com/huuthuan-nguyen/manabie/app/middleware"
	"github.com/huuthuan-nguyen/manabie/app/model"
	"github.com/huuthuan-nguyen/manabie/app/render"
	"github.com/huuthuan-nguyen/manabie/app/request"
	"github.com/huuthuan-nguyen/manabie/app/transformer"
	"github.com/huuthuan-nguyen/manabie/app/utils"
	"net/http"
	"time"
)

// UserRegister /**
func (handler *Handler) UserRegister(w http.ResponseWriter, r *http.Request) {
	credentialsRequest := request.Credentials{}
	userModel := model.User{}

	if err := credentialsRequest.Bind(r, &userModel); err != nil {
		render.Error(w, r, err)
		return
	}

	if err := userModel.Create(r.Context(), handler.db); err != nil {
		utils.PanicInternalServerError(err)
		return
	}

	userTransformer := &transformer.UserTransformer{}
	userItem := transformer.NewItem(userModel, userTransformer)

	// render user item to JSON
	render.JSON(w, r, userItem)
}

// UserLogin /**
func (handler *Handler) UserLogin(w http.ResponseWriter, r *http.Request) {
	credentialsRequest := request.Credentials{}
	userModel := model.User{}

	if err := credentialsRequest.Bind(r, &userModel); err != nil {
		render.Error(w, r, err)
		return
	}

	// find user by email in database
	currentUser, err := model.FindOneUserByEmail(r.Context(), userModel.Email, handler.db)
	if err != nil {
		render.Error(w, r, err)
	}

	// compare password
	if !model.CheckPasswordHash(userModel.Password, currentUser.Password) {
		// throw error
		utils.PanicUnauthorized()
		return
	}

	// expired at
	expirationTime := time.Now().Add(1440 * time.Minute) // one day expiration
	// create JWT claims with expiration and email
	claims := &middleware.Payload{
		Email: userModel.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // expired at in unix milliseconds format
		},
	}
	// create token with HS256 algorithm and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// token string
	tokenString, err := token.SignedString([]byte(handler.config.Server.Secret))
	if err != nil {
		utils.PanicInternalServerError(err)
		return
	}

	// return token
	tokenModel := model.Token{
		AccessToken: tokenString,
		ExpiredAt:   expirationTime,
	}
	tokenTransformer := &transformer.TokenTransformer{}
	tokenItem := transformer.NewItem(tokenModel, tokenTransformer)

	// render token item to JSON
	render.JSON(w, r, tokenItem)
}
