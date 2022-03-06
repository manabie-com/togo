package usecase

import (
	"fmt"

	"github.com/triet-truong/todo/domain"
	"github.com/triet-truong/todo/todo/dto"
	"github.com/triet-truong/todo/todo/model"
	"github.com/triet-truong/todo/todo/utils"
)

type TodoUseCase struct {
	repository domain.TodoRepository
	cacheStore domain.TodoCacheRepository
}

func NewTodoUseCase(repo domain.TodoRepository, cacheStore domain.TodoCacheRepository) *TodoUseCase {
	return &TodoUseCase{
		repository: repo,
		cacheStore: cacheStore,
	}
}

func (u *TodoUseCase) AddTodo(todoDto dto.TodoDto) error {
	cachedUser, err := u.cacheStore.GetCachedUser(todoDto.UserId)
	user := model.UserModel{}
	if err != nil {
		utils.ErrorLog(err)
		user, err = u.repository.SelectUser(todoDto.UserId)
		if err != nil {
			utils.ErrorLog(err)
			return fmt.Errorf("user not found")
		}
		cachedUser = model.UserRedisModel{
			ID:               user.ID,
			DailyRecordLimit: user.DailyRecordLimit,
			CurrentUsage:     0,
		}
	}

	cachedUser.CurrentUsage++
	if cachedUser.CurrentUsage > cachedUser.DailyRecordLimit {
		return fmt.Errorf("exceeded daily limit")
	}

	err = u.repository.InsertItem(model.TodoItemModel{
		Content: todoDto.Content,
		UserID:  todoDto.UserId,
		IsDone:  false,
	})
	if err != nil {
		utils.ErrorLog(err)
		return err
	}

	u.cacheStore.SetUser(cachedUser)
	return nil
}
