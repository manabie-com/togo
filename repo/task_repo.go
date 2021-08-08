package repo

import (
	"github.com/manabie-com/togo/common/context"
	"github.com/manabie-com/togo/domain/model"
)

type TaskRepository interface {
	Insert(ctx context.Context, task *model.Task) error
	FindTaskByUserIdAndDate(ctx context.Context, userId string, createdDate string) ([]*model.Task, error)
}

type task struct {
	ID          string
	UserID      string
	Content     string
	CreatedDate string
}

func (t *task) TableName() string {
	return "tasks"
}

func NewTaskRepo() TaskRepository {
	return &taskRepo{}
}

type taskRepo struct {
}

func (t *taskRepo) FindTaskByUserIdAndDate(ctx context.Context, userId string, createdDate string) ([]*model.Task, error) {
	var tasks []*task
	db := ctx.GetDb()
	err := db.Find(&tasks, "user_id = ? and created_date = ?", userId, createdDate).Error
	if err != nil {
		return nil, err
	}
	return mapTasksGormToModel(tasks), nil
}

func (t *taskRepo) Insert(ctx context.Context, task *model.Task) error {
	taskGorm := mapTaskModelToModel(task)
	db := ctx.GetDb()
	err := db.Save(taskGorm).Error
	if err != nil {
		return err
	}
	return nil
}

func mapTaskModelToModel(taskModel *model.Task) *task {
	return &task{
		ID:          taskModel.ID,
		Content:     taskModel.Content,
		UserID:      taskModel.UserID,
		CreatedDate: taskModel.CreatedDate,
	}
}

func mapTasksGormToModel(tasksGorm []*task) []*model.Task {
	tasksModel := make([]*model.Task, 0)
	for _, task := range tasksGorm {
		tasksModel = append(tasksModel, &model.Task{
			ID:          task.ID,
			Content:     task.Content,
			UserID:      task.UserID,
			CreatedDate: task.CreatedDate,
		})
	}
	return tasksModel
}
