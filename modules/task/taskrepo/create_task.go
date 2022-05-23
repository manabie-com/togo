package taskrepo

import (
	"context"
	"github.com/japananh/togo/common"
	"github.com/japananh/togo/modules/task/taskmodel"
	"github.com/japananh/togo/modules/user/usermodel"
)

type CreateTaskStore interface {
	FindTaskByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*taskmodel.Task, error)
	CountUserDailyTask(ctx context.Context, createdBy int) (int, error)
	CreateTask(ctx context.Context, data *taskmodel.TaskCreate) error
}

type UserStore interface {
	FindUser(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*usermodel.User, error)
}

type createTaskRepo struct {
	store     CreateTaskStore
	userStore UserStore
}

func NewCreateTaskRepo(store CreateTaskStore, userStore UserStore) *createTaskRepo {
	return &createTaskRepo{
		store:     store,
		userStore: userStore,
	}
}

func (biz *createTaskRepo) CreateTask(
	ctx context.Context,
	data *taskmodel.TaskCreate,
) error {
	// check for exited assignee
	if data.AssigneeId > 0 {
		if _, err := biz.userStore.FindUser(ctx, map[string]interface{}{"id": data.AssigneeId}); err != nil {
			return common.NewCustomError(nil, "invalid assignee", "ErrInvalidAssignee")
		}
	}

	// check for existed parent task
	if data.ParentId > 0 {
		_, err := biz.store.FindTaskByCondition(ctx, map[string]interface{}{"id": data.ParentId})
		if err != nil {
			return common.NewCustomError(nil, "invalid parent task", "ErrInvalidParentTask")
		}
	}

	// user cannot create more task than his/her daily task limit
	count, err := biz.store.CountUserDailyTask(ctx, data.CreatedBy)
	if err != nil {
		return common.ErrDB(err)
	}

	createdBy, err := biz.userStore.FindUser(ctx, map[string]interface{}{"id": data.CreatedBy})
	if err != nil {
		return common.NewCustomError(nil, "invalid task creator", "ErrInvalidCreatedBy")
	}

	if count >= createdBy.DailyTaskLimit {
		return common.NewCustomError(nil, "exceed daily task limit", "ErrExceedDailyTaskLimit")
	}

	if err := biz.store.CreateTask(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(taskmodel.EntityName, err)
	}

	return nil
}
