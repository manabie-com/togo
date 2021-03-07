package user

import (
	"togo/src"
	"togo/src/entity/user"
	gErrors "togo/src/infra/error"
	userRepo "togo/src/infra/repository/user"
	"togo/src/infra/service"
	"togo/src/schema"
)

type UserWorkflow struct {
	repository user.IUserRepository
	jwtService src.IJWTService
}

func (this *UserWorkflow) Login(data *schema.LoginRequest) (*schema.LoginResponse, error) {
	result, err := this.repository.FindOne(&user.User{
		ID:       data.UserId,
		Password: data.Password,
	})
	if err != nil {
		return nil, err
	}

	// Dump token for simple demo app (Real appliaction would store permission in database)
	permissions := []string{}

	if result.ID == "firstUser" {
		permissions = append(permissions, "task.create")
	}

	token, err := this.jwtService.CreateToken(&src.TokenData{
		UserId:      result.ID,
		Permissions: permissions,
	})
	if err != nil {
		return nil, gErrors.NewInternalServerError(src.CREATE_TOKEN_FAIL, err)
	}

	return &schema.LoginResponse{
		UserId: result.ID,
		Token:  token,
	}, nil
}

func (this *UserWorkflow) Register(data *schema.RegisterRequest) (*schema.RegisterResponse, error) {
	return nil, nil
}

func (this *UserWorkflow) CreateTaskByOwner(data *schema.CreateTaskByOwnerRequest) (*schema.CreateTaskByOwnerResponse, error) {
	return nil, nil
}

func (this *UserWorkflow) DeleteTaskByOwner(data *schema.DeleteTaskByOwnerRequest) (*schema.DeleteTaskByOwnerResponse, error) {
	return nil, nil
}

func NewUserWorkflow() IUserWorkflow {
	return &UserWorkflow{
		userRepo.NewUserRepository(),
		service.NewJWTService(),
	}
}
