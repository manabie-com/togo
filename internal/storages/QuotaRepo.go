package storages

import (
	"fmt"
	"log"
	"github.com/google/uuid"	
	"github.com/manabie-com/togo/internal/utils"
)

const QUOTA_LIMIT = 5

type IQuotaRepo interface {
	FindByUserIdFromRepo(user_id string) (Quota, error)	
	ReplaceFromRepo(quota Quota) error
	InitQuota(user_id string) Quota
}

type QuotaRepo struct {
	IDBHandler
}

func (quotaRepo *QuotaRepo) FindByUserIdFromRepo(user_id string) (Quota, error) {

	rows, err := quotaRepo.Query(fmt.Sprintf("SELECT id, quota, user_id, start_time FROM quotas WHERE user_id = '%s'", user_id))
	if err != nil {
		return Quota{}, err
	}	
	
	q := Quota{}
	for rows.Next() {
		err := rows.Scan(&q.ID, &q.Quota, &q.UserID, &q.StartTime)
		if err != nil {
			return Quota{}, err
		}
	}

	return q, nil		
}

func (quotaRepo *QuotaRepo) ReplaceFromRepo(q Quota) error {

	err := quotaRepo.Execute(fmt.Sprintf("REPLACE INTO quotas (id, quota, user_id, start_time) VALUES ('%s', '%d', '%s', '%s')", q.ID, q.Quota, q.UserID, q.StartTime))	
	if err != nil {
		log.Println(err)
		return err
	}

	return nil	
}

func (quotaRepo *QuotaRepo) InitQuota(user_id string) Quota {
	q := Quota{}
	q.ID = uuid.New().String()
	q.UserID = user_id
	q.Quota = QUOTA_LIMIT
	q.StartTime = utils.GetStartTime()

	return q
}