package services

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"time"

	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/redis"
)


func (s *ToDoService) listTasks(resp http.ResponseWriter, req *http.Request) {
	var isParamCreatedDate bool
	var cd time.Time
	id, _ := userIDFromCtx(req.Context())
	createdDate := req.FormValue("created_date")
	if len(createdDate) > 0 {
		isParamCreatedDate = true
		t, err := time.Parse("2006-01-02", createdDate)
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(resp).Encode(map[string]string{
				"error": "created_date is failed",
			})
			return
		}
		cd = t
	}
	resp.Header().Set("Content-Type", "application/json")

	tasks := storages.GetTasks()
	err := s.Store.Where("user_id = ? and (false = ? or created_at > ?) ", id, isParamCreatedDate, cd).Find(&tasks).Error
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "Get list Task Failed",
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string][]storages.Task{
		"data": tasks,
	})
}

func (s *ToDoService) addTask(resp http.ResponseWriter, req *http.Request) {
	ct := storages.GetCreateTask()
	err := json.NewDecoder(req.Body).Decode(&ct)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	resp.Header().Set("Content-Type", "application/json")

	t := storages.GetTask()
	now := time.Now().UTC()
	userID, _ := userIDFromCtx(req.Context())
	t.ID = uuid.New().String()
	t.UserID = userID
	t.CreatedAt = now
	t.UpdatedAt = now
	t.Content = ct.Content

	err = validationLimitedToCreate(req.Context(), &t, 6)
	if err != nil {
		resp.WriteHeader(http.StatusForbidden)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	err = s.Store.Create(&t).Error
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]*storages.Task{
		"data": &t,
	})
}

func userIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(userAuthKey(0))
	id, ok := v.(string)
	return id, ok
}

func validationLimitedToCreate(ctx context.Context, t *storages.Task, n int64) error {
	client := redis.Init()
	year, month, day := time.Now().Date()
	sYear := strconv.Itoa(year)
	sMonth := strconv.Itoa(int(month))
	sDay := strconv.Itoa(day)
	key := t.UserID+sYear+sMonth+sDay

	err := client.SetNX(ctx, key, 0)
	if err != nil {
		return err
	}

	integer, err := client.Incr(ctx, key)
	if err != nil {
		return err
	}
	if integer >= n {
		return errors.New("Create only 5 tasks only per day")
	}

	return nil
}
