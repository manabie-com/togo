package jwt

import (
	"errors"
	tokenprovider "github.com/manabie-com/togo/token_provider"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	ErrNotFound      = errors.New("tokenprovider not found")
	ErrEncodingToken = errors.New("error encoding the tokenprovider")
	ErrInvalidToken  = errors.New("invalid tokenprovider provided")
)

type jwtProvider struct {
	secret string
	expiry int
}

func NewTokenJWTProvider(secret string, expiry int) *jwtProvider {
	return &jwtProvider{secret: secret, expiry: expiry}
}

type myClaims struct {
	tokenprovider.JwtPayload `json:",inline"`
	jwt.StandardClaims
}

func (j *jwtProvider) Generate(id int, expiry int) (*tokenprovider.Token, error) {
	// set expiry time
	j.expiry = expiry
	// generate the JWT
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		tokenprovider.JwtPayload{
			UserId: id,
		},
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(j.expiry)).Unix(),
		},
	})

	myToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	// return the tokenprovider
	return &tokenprovider.Token{
		Token:   myToken,
		Expiry:  j.expiry,
		Created: time.Now(),
	}, nil
}

func (j *jwtProvider) GenRefreshToken(pl tokenprovider.IPayload, expiry int) (*tokenprovider.Token, error) {
	j.expiry = expiry
	// generate the JWT
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		tokenprovider.JwtPayload{
			UserId: pl.GetUserId(),
		},
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(j.expiry)).Unix(),
		},
	})

	myToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	// return the tokenprovider
	return &tokenprovider.Token{
		Token:   myToken,
		Expiry:  j.expiry,
		Created: time.Now(),
	}, nil
}

func (j *jwtProvider) GenAccessToken(pl tokenprovider.IPayload, expiry int) (*tokenprovider.Token, error) {
	// set expiry time
	j.expiry = expiry
	// generate the JWT
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		tokenprovider.JwtPayload{
			UserId: pl.GetUserId(),
		},
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(j.expiry)).Unix(),
		},
	})

	myToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	// return the tokenprovider
	return &tokenprovider.Token{
		Token:   myToken,
		Expiry:  j.expiry,
		Created: time.Now(),
	}, nil
}

func (j *jwtProvider) Validate(myToken string) (tokenprovider.IPayload, error) {
	// parse the public key
	res, err := jwt.ParseWithClaims(myToken, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, ErrInvalidToken
	}

	// validate the tokenprovider
	if !res.Valid {
		return nil, ErrInvalidToken
	}

	claims, ok := res.Claims.(*myClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	// return the tokenprovider
	return &claims.JwtPayload, nil
}

func (j *jwtProvider) String() string {
	return "JWT implement Provider"
}
