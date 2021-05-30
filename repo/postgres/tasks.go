package postgre

import (
	"fmt"

	"github.com/manabie-com/togo/db"
	"github.com/manabie-com/togo/modules/tasks"
)

type TasksRepo struct {
}

func (r TasksRepo) Create(req tasks.Tasks) (interface{}, error) {

	result := db.DB.Create(&req)
	if result.Error == nil {
		fmt.Println("Inserted!")
		return req, nil
	}
	return nil, result.Error

}

func (r TasksRepo) GetList(createdDate string) ([]tasks.Tasks, error) {
	var task []tasks.Tasks
	db.DB.Model(&tasks.Tasks{}).Where("created_date = ?", createdDate).Find(&task)
	return task, nil
}

func (r TasksRepo) CountTask(userId string, createdDate string) int64 {
	var count int64 = 0
	db.DB.Model(&tasks.Tasks{}).Where("user_id=? and created_date = ?", userId, createdDate).Count(&count)
	return count
}
