package usecase

import (
	"errors"
	e "lntvan166/togo/internal/entities"
	mockdb "lntvan166/togo/pkg/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var user = &e.User{
	ID:       1,
	Username: "user",
	Password: "admin",
	Plan:     "free",
	MaxTodo:  10,
}

var user2 = &e.User{
	ID:       2,
	Username: "user2",
	Password: "admin",
	Plan:     "vip",
	MaxTodo:  20,
}

var userWithPasswordHashed = &e.User{
	ID:       1,
	Username: "user",
	Password: "$2a$10$3JwrD1vVYd15jdbv9HAVw.cEBe9ou6YCS0ypTzlSjLC/h8qbLXW6m",
	Plan:     "free",
	MaxTodo:  10,
}

var userAfterUpgrade = &e.User{
	ID:       1,
	Username: "user",
	Password: "admin",
	Plan:     "vip",
	MaxTodo:  20,
}

var users = []*e.User{
	user,
	user2,
}

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mockdb.NewMockUserRepository(ctrl)
	taskRepo := mockdb.NewMockTaskRepository(ctrl)
	crypto := mockdb.NewMockAppCrypto(ctrl)

	userUsecase := NewUserUsecase(userRepo, taskRepo, crypto)

	// before test
	userRepo.EXPECT().AddUser(user).Return(nil)
	userRepo.EXPECT().GetUserByName(user.Username).Return(nil, errors.New("not found"))
	crypto.EXPECT().HashPassword(user.Password).Return(userWithPasswordHashed.Password)

	err := userUsecase.Register(user)
	assert.NoError(t, err)
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mockdb.NewMockUserRepository(ctrl)
	taskRepo := mockdb.NewMockTaskRepository(ctrl)
	crypto := mockdb.NewMockAppCrypto(ctrl)

	userUsecase := NewUserUsecase(userRepo, taskRepo, crypto)

	// before test
	userRepo.EXPECT().GetUserByName(user.Username).Return(userWithPasswordHashed, nil).AnyTimes()
	crypto.EXPECT().ComparePassword(user.Password, userWithPasswordHashed.Password).Return(true).AnyTimes()

	token, err := userUsecase.Login(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestGetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mockdb.NewMockUserRepository(ctrl)
	taskRepo := mockdb.NewMockTaskRepository(ctrl)
	crypto := mockdb.NewMockAppCrypto(ctrl)

	userUsecase := NewUserUsecase(userRepo, taskRepo, crypto)

	// before test
	userRepo.EXPECT().GetAllUsers().Return(users, nil)

	usersDB, err := userUsecase.GetAllUsers()
	assert.NoError(t, err)
	assert.Equal(t, users, usersDB)
}

func TestGetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mockdb.NewMockUserRepository(ctrl)
	taskRepo := mockdb.NewMockTaskRepository(ctrl)
	crypto := mockdb.NewMockAppCrypto(ctrl)

	userUsecase := NewUserUsecase(userRepo, taskRepo, crypto)

	// before test
	userRepo.EXPECT().GetUserByID(user.ID).Return(user, nil)

	userDB, err := userUsecase.GetUserByID(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user, userDB)
}

func TestGetUserByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mockdb.NewMockUserRepository(ctrl)
	taskRepo := mockdb.NewMockTaskRepository(ctrl)
	crypto := mockdb.NewMockAppCrypto(ctrl)

	userUsecase := NewUserUsecase(userRepo, taskRepo, crypto)

	// before test
	userRepo.EXPECT().GetUserByName(user.Username).Return(user, nil)

	userDB, err := userUsecase.GetUserByName(user.Username)
	assert.NoError(t, err)
	assert.Equal(t, user, userDB)
}

func TestGetUserIDByUsername(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mockdb.NewMockUserRepository(ctrl)
	taskRepo := mockdb.NewMockTaskRepository(ctrl)
	crypto := mockdb.NewMockAppCrypto(ctrl)

	userUsecase := NewUserUsecase(userRepo, taskRepo, crypto)

	// before test
	userRepo.EXPECT().GetUserByName(user.Username).Return(user, nil)

	id, err := userUsecase.GetUserIDByUsername(user.Username)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, id)
}

func TestGetMaxTaskByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mockdb.NewMockUserRepository(ctrl)
	taskRepo := mockdb.NewMockTaskRepository(ctrl)
	crypto := mockdb.NewMockAppCrypto(ctrl)

	userUsecase := NewUserUsecase(userRepo, taskRepo, crypto)

	// before test
	userRepo.EXPECT().GetUserByID(user.ID).Return(user, nil)

	maxTodo, err := userUsecase.GetMaxTaskByUserID(user.ID)
	assert.NoError(t, err)
	assert.EqualValues(t, user.MaxTodo, maxTodo)
}

func TestGetPlan(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mockdb.NewMockUserRepository(ctrl)
	taskRepo := mockdb.NewMockTaskRepository(ctrl)
	crypto := mockdb.NewMockAppCrypto(ctrl)

	userUsecase := NewUserUsecase(userRepo, taskRepo, crypto)

	// before test
	userRepo.EXPECT().GetUserByName(user.Username).Return(user, nil)

	plan, err := userUsecase.GetPlan(user.Username)
	assert.NoError(t, err)
	assert.Equal(t, user.Plan, plan)
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mockdb.NewMockUserRepository(ctrl)
	taskRepo := mockdb.NewMockTaskRepository(ctrl)
	crypto := mockdb.NewMockAppCrypto(ctrl)

	userUsecase := NewUserUsecase(userRepo, taskRepo, crypto)

	// before test
	userRepo.EXPECT().UpdateUser(user).Return(nil)

	err := userUsecase.UpdateUser(user)
	assert.NoError(t, err)
}

func TestUpgradePlan(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mockdb.NewMockUserRepository(ctrl)
	taskRepo := mockdb.NewMockTaskRepository(ctrl)
	crypto := mockdb.NewMockAppCrypto(ctrl)

	userUsecase := NewUserUsecase(userRepo, taskRepo, crypto)

	// before test
	userRepo.EXPECT().GetUserByID(user.ID).Return(user, nil)
	userRepo.EXPECT().UpdateUser(user).Return(nil)

	err := userUsecase.UpgradePlan(user.ID, userAfterUpgrade.Plan, int(userAfterUpgrade.MaxTodo))
	assert.NoError(t, err)
}

func TestCheckUserExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mockdb.NewMockUserRepository(ctrl)
	taskRepo := mockdb.NewMockTaskRepository(ctrl)
	crypto := mockdb.NewMockAppCrypto(ctrl)

	userUsecase := NewUserUsecase(userRepo, taskRepo, crypto)

	// before test
	userRepo.EXPECT().GetUserByName(user.Username).Return(user, nil)

	isExist := userUsecase.CheckUserExist(user.Username)
	assert.True(t, isExist)
}

func TestDeleteUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mockdb.NewMockUserRepository(ctrl)
	taskRepo := mockdb.NewMockTaskRepository(ctrl)
	crypto := mockdb.NewMockAppCrypto(ctrl)

	userUsecase := NewUserUsecase(userRepo, taskRepo, crypto)

	// before test
	taskRepo.EXPECT().DeleteAllTaskOfUser(user.ID).Return(nil)
	userRepo.EXPECT().DeleteUserByID(user.ID).Return(nil)

	err := userUsecase.DeleteUserByID(user.ID)
	assert.NoError(t, err)
}
