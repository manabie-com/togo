package togo_appgrpc

import (
	"context"
	"encoding/json"
	"flag"

	"github.com/phathdt/libs/go-sdk/logger"
	"github.com/phathdt/libs/go-sdk/sdkcm"
	protos "github.com/phathdt/libs/togo_proto/out/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient interface {
	FindUser(ctx context.Context, conditions map[string]interface{}) (*sdkcm.SimpleUser, error)
}

type userClient struct {
	prefix      string
	url         string
	gwSupported bool
	gwPort      int
	logger      logger.Logger
	client      protos.UserServiceClient
}

func NewUserClient(prefix string) *userClient {
	return &userClient{
		prefix: prefix,
	}
}

func (uc *userClient) GetPrefix() string {
	return uc.prefix
}

func (uc *userClient) Get() interface{} {
	return uc
}

func (uc *userClient) Name() string {
	return uc.prefix
}

func (uc *userClient) InitFlags() {
	flag.StringVar(&uc.url, uc.GetPrefix()+"-url", "localhost:50051", "URL connect to grpc server")
}

func (uc *userClient) Configure() error {
	uc.logger = logger.GetCurrent().GetLogger(uc.prefix)

	uc.logger.Infoln("Setup user client service:", uc.prefix)

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())

	cc, err := grpc.Dial(uc.url, opts)

	if err != nil {
		return err
	}

	uc.client = protos.NewUserServiceClient(cc)

	uc.logger.Infof("Setup user client service success with url %s", uc.url)

	return nil
}

func (uc *userClient) Run() error {
	return uc.Configure()
}

func (uc *userClient) Stop() <-chan bool {
	c := make(chan bool)

	go func() {
		c <- true
		uc.logger.Infoln("Stopped")
	}()
	return c
}

func (uc *userClient) FindUser(ctx context.Context, conditions map[string]interface{}) (*sdkcm.SimpleUser, error) {
	bytes, _ := json.Marshal(conditions)
	rs, err := uc.client.FindUser(ctx, &protos.FindUserRequest{Cond: string(bytes)})
	if err != nil {
		return nil, sdkcm.ErrCannotGetEntity("user", err)
	}

	user := sdkcm.SimpleUser{
		SQLModel: *sdkcm.NewUpsertSQLModel(int(rs.User.Id)),
		Email:    rs.User.Email,
	}

	return &user, nil
}
