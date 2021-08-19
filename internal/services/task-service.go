package services

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/models"
	repository "github.com/manabie-com/togo/internal/repositories"
	"github.com/manabie-com/togo/internal/utils"
)

// ToDoService implement HTTP server
type ToDoTaskService struct {
	JWTKey string
	Store  *repository.DB
}

func NewToDoTaskService(db *repository.DB, jwtKey string) *ToDoTaskService {
	return &ToDoTaskService{
		JWTKey: jwtKey,
		Store:  db,
	}
}

func (t *ToDoTaskService) ListTasks(ctx context.Context, id string, createdDate string) ([]*models.Task, error) {
	return t.Store.RetrieveTasks(
		ctx,
		sql.NullString{
			String: id,
			Valid:  true,
		},
		sql.NullString{
			String: createdDate,
			Valid:  true,
		},
	)
}

func (s *ToDoTaskService) AddTask(ctx context.Context, addingTask *models.Task) (*models.Task, error) {
	now := time.Now()
	userID, _ := utils.NewJwtUtil(s.JWTKey).UerIDFromCtx(ctx)
	createdDate := now.Format("2006-01-02")
	addingTask.ID = uuid.New().String()
	addingTask.UserID = userID
	addingTask.CreatedDate = createdDate

	if isReachedLimitedTasksADay(ctx, s.Store, userID, createdDate) {
		return nil, errors.New("Status Not Acceptable")
	}

	err := s.Store.AddTask(ctx, addingTask)
	if err != nil {
		return nil, err
	}

	return addingTask, nil
}

func isReachedLimitedTasksADay(ctx context.Context, db *repository.DB, userId string, createdDate string) bool {
	sqlUserIdVal := sql.NullString{
		String: userId,
		Valid:  true,
	}
	sqlCreatedDateVal := sql.NullString{
		String: createdDate,
		Valid:  true,
	}

	tasks, err := db.RetrieveTasks(
		ctx,
		sqlUserIdVal,
		sqlCreatedDateVal,
	)

	if err != nil {
		log.Println(err)

		return true
	}

	if len(tasks) >= db.RetrieveMaxToDoSetting(ctx, sqlUserIdVal) {
		return true
	}

	return false
}
