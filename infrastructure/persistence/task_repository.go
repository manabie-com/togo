package persistence

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jfzam/togo/domain/entity"
	"github.com/jfzam/togo/domain/repository"
	"gorm.io/gorm"
)

type TaskRepo struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepo {
	return &TaskRepo{db}
}

//TaskRepo implements the repository.TaskRepository interface
var _ repository.TaskRepository = &TaskRepo{}

func (r *TaskRepo) SaveTask(task *entity.Task, userLimit int64) (*entity.Task, map[string]string) {
	dbErr := map[string]string{}

	// check if has reached limit
	var tasks []entity.Task
	err := r.db.Debug().Where("created_at > ? and user_id = ?", time.Now().Format("02/01/2006"), task.UserID).Find(&tasks).Error
	if err != nil {
		dbErr["db_error"] = "database error; error in looking up for user task for today"
		return nil, dbErr
	}

	fmt.Print("userLimit:", userLimit, "tasks:", len(tasks))

	// check if user reached its task limit
	if int(userLimit) <= len(tasks) {
		fmt.Println("user reached its task limit for today")
		dbErr["user_limit"] = "user reached its task limit for today"
		return nil, dbErr
	}

	err = r.db.Debug().Create(&task).Error
	if err != nil {
		//since our title is unique
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_title"] = "task title already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return task, nil
}

func (r *TaskRepo) GetTask(id uint64) (*entity.Task, error) {
	var task entity.Task
	err := r.db.Debug().Where("id = ?", id).Take(&task).Error
	if err != nil {
		return nil, errors.New("database error, please try again")
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("food not found")
	}
	return &task, nil
}

func (r *TaskRepo) GetAllTask() ([]entity.Task, error) {
	var tasks []entity.Task
	err := r.db.Debug().Limit(100).Order("created_at desc").Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	return tasks, nil
}
