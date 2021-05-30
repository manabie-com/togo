package tasks

import (
	"time"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/modules/auth"
	"github.com/manabie-com/togo/modules/common"
)

type Service struct {
	Repo TasksRepository
}

func (r Service) Create(req Tasks, userAuth auth.UserAuth) (interface{}, *common.RequestError) {

	now := time.Now()
	req.ID = uuid.New().String()
	req.UserId = userAuth.UserID
	req.CreatedDate = now.Format("2006-01-02")
	countTask := r.Repo.CountTask(req.UserId, req.CreatedDate)

	if countTask < int64(userAuth.MaxTodo) {

		task, err := r.Repo.Create(req)

		if err != nil {
			return nil, common.NewInternalError()
		}
		return task, nil
	} else {
		return nil, common.NewPermissionError("Tasks over limited today")
	}

}

func (r Service) GetList(phone string) ([]Tasks, *common.RequestError) {
	list, err := r.Repo.GetList(phone)

	if err != nil {
		return list, common.NewInternalError()
	}
	return list, nil
}
