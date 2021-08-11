package domain

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/common"
	"github.com/manabie-com/togo/common/context"
	"github.com/manabie-com/togo/common/middleware"
	"github.com/manabie-com/togo/domain/model"
	"github.com/manabie-com/togo/repo"
	"log"
	"time"
)

type TodoService interface {
	GetAccessToken(ctx context.Context, username, password string) (string, error)
	CreateTask(ctx context.Context, content string) (*model.Task, error)
	GetTaskAtDate(ctx context.Context, createDate time.Time) ([]*model.Task, error)
}

func NewTodoService(
	jwtKey string,
	tokenTimeout int,
	userRepo repo.UserRepository,
	taskRepo repo.TaskRepository,
) TodoService {
	return &todoService{
		jwtKey:       jwtKey,
		userRepo:     userRepo,
		tokenTimeout: tokenTimeout,
		taskRepo:     taskRepo,
	}
}

type todoService struct {
	jwtKey       string
	tokenTimeout int
	userRepo     repo.UserRepository
	taskRepo     repo.TaskRepository
}

func (t *todoService) GetTaskAtDate(ctx context.Context, createDate time.Time) ([]*model.Task, error) {
	log.Println("get task")
	userId := ctx.GetUserId()
	tasks, err := t.taskRepo.FindTaskByUserIdAndDate(ctx, userId, createDate)
	if err != nil {
		log.Println("can't find tasks: ", err)
		return nil, err
	}
	return tasks, nil
}

func (t *todoService) CreateTask(ctx context.Context, content string) (*model.Task, error) {
	log.Println("create task")
	userId := ctx.GetUserId()
	user, err := t.userRepo.GetUserById(ctx, userId)
	if err != nil {
		log.Println("error when get user: ", err)
		return nil, err
	}
	if user == nil {
		log.Println("can't found user")
		return nil, ErrUserNotFound
	}
	nowRounded := common.GetCurrentDateRounded()
	existingTasks, err := t.taskRepo.FindTaskByUserIdAndDate(ctx, userId, nowRounded)
	if err != nil {
		log.Println("can't find tasks")
		return nil, err
	}
	if len(existingTasks) >= user.MaxTodo {
		log.Println("reach limit tasks in a day")
		return nil, ErrFailPrecondition
	}

	task := &model.Task{
		ID:          uuid.New().String(),
		Content:     content,
		UserID:      userId,
		CreatedDate: nowRounded,
	}
	err = t.taskRepo.Insert(ctx, task)
	if err != nil {
		log.Println("can't add task")
		return nil, err
	}
	return task, nil
}

func (t *todoService) GetAccessToken(ctx context.Context, username, password string) (string, error) {
	log.Println("get access token")
	user, err := t.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		log.Println("error when get user: ", err)
		return "", err
	}
	if user == nil {
		log.Println("can't found user with username: ", username)
		return "", ErrUserNotFound
	}
	if user.Password != password {
		log.Println("password not match")
		return "", ErrFailPrecondition
	}
	return t.createToken(user.Id)
}

func (t *todoService) createToken(id string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims[middleware.UserIdClaimName] = id
	atClaims["exp"] = time.Now().Add(time.Minute * time.Duration(t.tokenTimeout)).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(t.jwtKey))
	if err != nil {
		return "", err
	}
	return token, nil
}
