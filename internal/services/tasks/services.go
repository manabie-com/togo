package tasks

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/utils"
)

// ToDoService implement HTTP server
type ToDoService struct {
	repo storages.Repository
}

func SetupNewService(r storages.Repository) *ToDoService {
	return &ToDoService{repo: r}
}

func (s *ToDoService) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL.Path)
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "*")

	if req.Method == http.MethodOptions {
		resp.WriteHeader(http.StatusOK)
		return
	}
}

func (s *ToDoService) ListTasks(context context.Context, id string, created_date sql.NullString) ([]*storages.Task, error) {
	tasks, err := s.repo.RetrieveTasks(
		context,
		sql.NullString{
			String: id,
			Valid:  true,
		},
		created_date,
	)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *ToDoService) AddTask(ctx context.Context, id sql.NullString, t storages.Task) (*storages.Task, error) {
	createdDateInSqlNullString := utils.ConvertStringToSqlNullString(time.Now().Format("2006-01-02"))

	user, err := s.repo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	tasks, err := s.repo.RetrieveTasks(
		ctx,
		id,
		createdDateInSqlNullString,
	)
	if err != nil {
		return nil, err
	}
	if len(tasks) == int(user.MaxTodo) {
		return nil, errors.New("exceed today maximum allowed number of tasks")
	}

	t.UserID = id.String
	t.CreatedDate = createdDateInSqlNullString.String

	taskId, err := s.repo.AddTask(ctx, &t)
	if err != nil {
		return nil, err
	}

	t.ID = taskId

	return &t, nil
}

func (s *ToDoService) DeleteTaskByDate(ctx context.Context, userID sql.NullString, createdDate sql.NullString) error {
	err := s.repo.DeleteTaskByDate(
		ctx,
		userID,
		createdDate,
	)
	return err
}

type userAuthKey int8

func (s *ToDoService) UserIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(userAuthKey(0))
	id, ok := v.(string)
	return id, ok
}
