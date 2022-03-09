package usecase

import (
	"fmt"

	"github.com/triet-truong/todo/todo"
	"github.com/triet-truong/todo/todo/dto"
	"github.com/triet-truong/todo/todo/model"
	"github.com/triet-truong/todo/utils"
)

type TodoUseCase struct {
	repository todo.TodoRepository
	cacheStore todo.TodoCacheRepository
}

func NewTodoUseCase(repo todo.TodoRepository, cacheStore todo.TodoCacheRepository) *TodoUseCase {
	return &TodoUseCase{
		repository: repo,
		cacheStore: cacheStore,
	}
}

func (u *TodoUseCase) AddTodo(todoDto dto.TodoDto) error {
	cachedUser, err := u.getCachedUser(todoDto)
	if err != nil {
		return err
	}

	cachedUser.CurrentUsage++
	if cachedUser.CurrentUsage > cachedUser.DailyRecordLimit {
		return fmt.Errorf("exceeded daily limit")
	}

	if err = u.repository.InsertItem(model.TodoItemModel{
		Content: todoDto.Content,
		UserID:  todoDto.UserId,
		IsDone:  false,
	}); err != nil {
		utils.ErrorLog(err)
		return err
	}

	u.cacheStore.SetUser(cachedUser)
	return nil
}

func (u *TodoUseCase) getCachedUser(todoDto dto.TodoDto) (model.UserRedisModel, error) {
	cachedUser, err := u.cacheStore.GetCachedUser(todoDto.UserId)
	user := model.UserModel{}
	if err != nil {
		utils.ErrorLog(err)
		user, err = u.repository.SelectUser(todoDto.UserId)
		if err != nil {
			utils.ErrorLog(err)
			return model.UserRedisModel{}, fmt.Errorf("user not found")
		}
		cachedUser = model.UserRedisModel{
			ID:               user.ID,
			DailyRecordLimit: user.DailyRecordLimit,
			CurrentUsage:     0,
		}
	}
	return cachedUser, nil
}
