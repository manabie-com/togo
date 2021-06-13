package services

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/model"
	"github.com/manabie-com/togo/internal/storages/ent"
	"github.com/manabie-com/togo/internal/storages/ent/user"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService interface {
	Login(ctx context.Context, credential model.LoginCredential) (*model.AccessToken, error)

	CreateUser(ctx context.Context, username string, pwd string) (*ent.User, error)
}

type authServiceImpl struct {
	jwtKey string
	client *ent.Client
}

func NewAuthService(jwtKey string, client *ent.Client) AuthService {

	auth := &authServiceImpl{jwtKey: jwtKey, client: client}
	auth.initDb()
	return auth
}

func (a *authServiceImpl) Login(ctx context.Context, credential model.LoginCredential) (*model.AccessToken, error) {
	foundUser, err := a.client.User.Query().Where(user.UsernameEQ(credential.UserName)).Only(ctx)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(credential.Password))
	if err != nil {
		return nil, err
	}

	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = foundUser.UserID
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(a.jwtKey))
	if err != nil {
		return nil, err
	}
	return &model.AccessToken{Token: token}, nil

}

func (a *authServiceImpl) CreateUser(ctx context.Context, username string, pwd string) (*ent.User, error) {
	cipherPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)

	u, err := a.client.User.
		Create().
		SetUserID(uuid.NewString()).
		SetUsername(username).
		SetPassword(string(cipherPwd)).
		Save(ctx)
	return u, err
}

func (a *authServiceImpl) initDb() {
	ctx := context.Background()
	if err := a.client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	_, err := a.CreateUser(ctx, "firstUser", "example")
	if err == nil {
		log.Info("created default user")
	} else {
		//TODO: check user already existed
		log.WithFields(log.Fields{
			"error": err,
		}).Info("could not create default user")
	}
}
