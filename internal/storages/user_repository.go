package storages

import (
	"context"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/storages/ent"
	"github.com/manabie-com/togo/internal/storages/ent/user"
	log "github.com/sirupsen/logrus"
)

type UserRepository interface {
	FindByUsername(ctx context.Context, username string) (*ent.User, error)

	FindByUserId(ctx context.Context, userId string) (*ent.User, error)

	CreateUser(ctx context.Context, username string, pwd string) (*ent.User, error)

	InitDb()
}

type userRepositoryImpl struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) UserRepository {
	return &userRepositoryImpl{client: client}
}

func (t *userRepositoryImpl) FindByUsername(ctx context.Context, username string) (*ent.User, error) {
	return t.client.User.Query().Where(user.UsernameEQ(username)).Only(ctx)

}

func (t *userRepositoryImpl) FindByUserId(ctx context.Context, userId string) (*ent.User, error) {
	return t.client.User.Query().Where(user.UserIDEQ(userId)).Only(ctx)
}

func (t *userRepositoryImpl) CreateUser(ctx context.Context, username string, pwd string) (*ent.User, error) {
	return t.client.User.
		Create().
		SetUserID(uuid.NewString()).
		SetUsername(username).
		SetPassword(pwd).
		Save(ctx)
}

func (t *userRepositoryImpl) InitDb() {
	ctx := context.Background()
	if err := t.client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	username := "firstUser"
	existedUser, _ := t.FindByUsername(ctx, username)
	if existedUser == nil {
		_, err := t.CreateUser(ctx, username, "example")
		if err == nil {
			log.Info("created default user")
		} else {
			log.WithFields(log.Fields{
				"error": err,
			}).Info("could not create default user")
		}
	}

}
