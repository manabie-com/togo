package handlers

import (
	"fmt"

	"example.com/m/v2/constants"
	authRepo "example.com/m/v2/internal/repositories/auth"
	taskRepo "example.com/m/v2/internal/repositories/task"
	userRepo "example.com/m/v2/internal/repositories/user"
	"example.com/m/v2/utils"

	authService "example.com/m/v2/internal/usecases/auth"
	taskService "example.com/m/v2/internal/usecases/task"
	userService "example.com/m/v2/internal/usecases/user"

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

var MainUC = &MainUseCase{}

func SetMainUseCase(muc *MainUseCase) {
	MainUC = muc
}

func NewUseCase(db *gorm.DB) MainUseCase {
	// Create repositories.
	authRepository := authRepo.NewAuthRepository(db)
	userRepository := userRepo.NewUserRepository(db)
	taskRepository := taskRepo.NewTaskRepository(db)

	//Create UseCase
	authUseCase := authService.NewAuthUseCase(authRepository)
	userUseCase := userService.NewUserUseCase(userRepository)
	taskUseCase := taskService.NewTaskUseCase(taskRepository)

	return MainUseCase{
		Auth: authUseCase,
		User: userUseCase,
		Task: taskUseCase,
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
