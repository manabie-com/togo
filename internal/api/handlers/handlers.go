package handlers

import (
	"fmt"

	"github.com/manabie-com/togo/constants"
	authRepo "github.com/manabie-com/togo/internal/repositories/auth"
	taskRepo "github.com/manabie-com/togo/internal/repositories/task"
	userRepo "github.com/manabie-com/togo/internal/repositories/user"
	"github.com/manabie-com/togo/utils"

	authService "github.com/manabie-com/togo/internal/usecases/auth"
	taskService "github.com/manabie-com/togo/internal/usecases/task"
	userService "github.com/manabie-com/togo/internal/usecases/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type MainUseCase struct {
	Auth authService.AuthUseCase
	User userService.UserUseCase
	Task taskService.TaskUseCase
}

type MainRepository struct {
	Auth authService.AuthRepository
	User userService.UserRepository
	Task taskService.TaskRepository
}

var MainUC = &MainUseCase{}

func SetMainUseCase(muc *MainUseCase) {
	MainUC = muc
}

func NewUseCase(db *gorm.DB) MainUseCase {
	// Create repositories.
	mainRepositories := NewRepositories(db)

	//Create UseCase
	authUseCase := authService.NewAuthUseCase(mainRepositories.Auth)
	userUseCase := userService.NewUserUseCase(mainRepositories.User)
	taskUseCase := taskService.NewTaskUseCase(mainRepositories.Task)

	return MainUseCase{
		Auth: authUseCase,
		User: userUseCase,
		Task: taskUseCase,
	}
}

func HandleService(db *gorm.DB) MainUseCase {
	return NewUseCase(db)
}

func NewRepositories(db *gorm.DB) *MainRepository {
	return &MainRepository{
		Auth: authRepo.NewAuthRepository(db),
		User: userRepo.NewUserRepository(db),
		Task: taskRepo.NewTaskRepository(db),
	}
}

type UserInfoFromCtx string
type UserInfoFromToken struct {
	ID            string
	MaxTaskPerDay string
}

func GetValueCookieFromCtx(ctx *gin.Context, keyCookie string) (*string, error) {
	cookie, err := ctx.Request.Cookie(keyCookie)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Don't have cookie with Key: %s", keyCookie))
	}

	if cookie == nil {
		return nil, errors.New("Cookie is nil")
	}

	return utils.String(cookie.Value), nil
}

func GetUserInfoFromToken(ctx *gin.Context) (*UserInfoFromToken, error) {
	tokenStr, err := GetValueCookieFromCtx(ctx, constants.CookieTokenKey)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Don't have cookie with Key1: %s", constants.CookieTokenKey))
	}

	if utils.SafeString(tokenStr) == "" {
		return nil, errors.New(fmt.Sprintf("Don't have cookie with Key2: %s", constants.CookieTokenKey))
	}

	claims := make(jwt.MapClaims)
	tkn, err := jwt.ParseWithClaims(utils.SafeString(tokenStr), claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.Env.JwtKey), nil
	})

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Don't have cookie with Key3: %s", constants.CookieTokenKey))
	}

	if tkn == nil || !tkn.Valid {
		return nil, errors.New(fmt.Sprintf("Don't have cookie with Key4: %s", constants.CookieTokenKey))
	}

	id, ok := claims["user_id"].(string)
	if !ok {
		return nil, errors.New(fmt.Sprintf("Don't have cookie with Key5: %s", constants.CookieTokenKey))
	}

	maxTaskPerDay, ok := claims["max_task_per_day"].(string)
	if !ok {
		return nil, errors.New(fmt.Sprintf("Don't have cookie with Key6: %s", constants.CookieTokenKey))
	}

	userInfo := &UserInfoFromToken{
		ID:            id,
		MaxTaskPerDay: maxTaskPerDay,
	}

	return userInfo, nil
}
