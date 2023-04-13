package jwt

import (
	"flag"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/phathdt/libs/go-sdk/plugin/tokenprovider"
)

type jwtProvider struct {
	name   string
	secret string
}

func NewJWTProvider(name string) *jwtProvider {
	return &jwtProvider{name: name}
}

func (p *jwtProvider) GetPrefix() string {
	return p.Name()
}

func (p *jwtProvider) Get() interface{} {
	return p
}

func (p *jwtProvider) Name() string {
	return p.name
}

func (p *jwtProvider) InitFlags() {
	flag.StringVar(&p.secret, "jwt-secret", "200Lab.io", "Secret key for generating JWT")
}

func (p *jwtProvider) Configure() error {
	return nil
}

func (p *jwtProvider) Run() error {
	return nil
}

func (p *jwtProvider) Stop() <-chan bool {
	c := make(chan bool)
	go func() {
		c <- true
	}()
	return c
}

func (j *jwtProvider) SecretKey() string {
	return j.secret
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayload, expiry int) (tokenprovider.Token, error) {
	// generate the JWT
	now := time.Now()

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		TokenPayload{
			UId: data.UserId(),
		},
		jwt.StandardClaims{
			ExpiresAt: now.Local().Add(time.Second * time.Duration(expiry)).Unix(),
			IssuedAt:  now.Local().Unix(),
			Id:        fmt.Sprintf("%d", now.UnixNano()),
		},
	})

	myToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	// return the token
	return &token{
		Token:   myToken,
		Expiry:  expiry,
		Created: now,
	}, nil
}

func (j *jwtProvider) Validate(myToken string) (tokenprovider.TokenPayload, error) {
	res, err := jwt.ParseWithClaims(myToken, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, tokenprovider.ErrInvalidToken
	}

	// validate the token
	if !res.Valid {
		return nil, tokenprovider.ErrInvalidToken
	}

	claims, ok := res.Claims.(*myClaims)

	if !ok {
		return nil, tokenprovider.ErrInvalidToken
	}

	// return the token
	return claims.Payload, nil
}

type myClaims struct {
	Payload tokenprovider.TokenPayload `json:"payload"`
	jwt.StandardClaims
}

type token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

func (t *token) GetToken() string {
	return t.Token
}

type TokenPayload struct {
	UId int `json:"user_id"`
}

func (p TokenPayload) UserId() int {
	return p.UId
}
