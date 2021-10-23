package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/storages"
	pg "github.com/manabie-com/togo/internal/storages/pg"
)

// ToDoService implement HTTP server
type ToDoService struct {
	Store  *pg.PgDB
}

func (s *ToDoService) ValidateUser(ctx context.Context, userID, pwd string) bool {
	return s.Store.ValidateUser(ctx, 
	sql.NullString{
		String: userID,
		Valid:  true,
	}, sql.NullString{
		String: pwd,
		Valid:  true,
	})
}


func (s *ToDoService) ListTasks(ctx context.Context, userID, createdDate string) ([]Task, error){
	tasks := []Task{}
	data, err := s.Store.RetrieveTasks(
		ctx,
		sql.NullString{
			String: userID,
			Valid:  true,
		},
		sql.NullString{
			String: createdDate,
			Valid:  true,
		},
	)
	if err != nil {
		return tasks, err
	}

	for _, task := range data {
		tasks = append(tasks, Task{
			ID         :task.ID,
			Content    : task.Content,
			UserID      : task.UserID,
			CreatedDate : task.CreatedDate,
		})
	}

	return tasks, nil
}

func (s *ToDoService) CountTodayTasks(ctx context.Context, userID string) (int, error) {
	return s.Store.CountUserTasks(ctx,
	sql.NullString{
		String: userID,
		Valid:  true,
	}, 
	sql.NullString{
		String: time.Now().Format("2006-01-02"),
		Valid:  true,
	})
}

func (s *ToDoService) AddTask(ctx context.Context, userID, content string) (Task, error) {
	t := Task{}
	t.ID = uuid.New().String()
	t.UserID = userID
	t.CreatedDate = time.Now().Format("2006-01-02")
	t.Content = content

	return t, s.Store.AddTask(ctx, &storages.Task{
		ID: t.ID,
		UserID: t.UserID,
		Content: t.Content,
		CreatedDate: t.CreatedDate,

	})
}

