package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/manabie-com/togo/domains"
	"time"
)

func NewTaskRepositoryImpl(db *sql.DB) domains.TaskRepository {
	return TaskRepositoryPostgresImpl{
		db: db,
	}
}

type TaskRepositoryPostgresImpl struct {
	db *sql.DB
}

func (t TaskRepositoryPostgresImpl) GetCountCreatedTaskTodayByUser(ctx context.Context, userId int64) (int64, error) {
	var query = `
		SELECT COUNT(*) as count
		FROM 
			tasks
		WHERE 
			user_id = $1 AND created_date::date = $2::date
		GROUP BY user_id
	`

	dateStr := time.Now().Format("2006-01-02")

	var count int64
	err := t.db.QueryRowContext(ctx, query, userId, dateStr).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, domains.ErrorNotFound
		}
		return 0, err
	}

	return count, err
}

func (t TaskRepositoryPostgresImpl) CreateTask(ctx context.Context, task *domains.Task) (*domains.Task, error) {
	stmt := `INSERT INTO tasks (content, user_id, created_date) VALUES ($1, $2, $3) RETURNING id`

	var lastId int64
	err := t.db.QueryRowContext(ctx, stmt, task.Content, task.UserId, time.Now()).Scan(&lastId)
	if err != nil {
		return nil, err
	}

	return t.GetTaskById(ctx, &domains.TaskByIdRequest{
		Id:     lastId,
		UserId: task.UserId,
	})
}

func (t TaskRepositoryPostgresImpl) GetTasks(ctx context.Context, request *domains.TaskRequest) ([]*domains.Task, error) {
	var query = `
		SELECT "id", "content", "user_id", "created_date" 
		FROM 
			tasks
		WHERE 
		%s
	`
	count := 0
	conditions := "1=1"
	params := make([]interface{}, 0)

	if request.UserId > 0 {
		count++
		params = append(params, request.UserId)
		conditions = conditions + " AND user_id = " + fmt.Sprintf("$%d", count)
	}

	if !request.CreatedDate.IsZero() {
		dateStr := request.CreatedDate.Format("2006-01-02")
		count++
		params = append(params, dateStr)
		conditions = conditions + " AND created_date::date = " + fmt.Sprintf("$%d::date", count)
	}

	tasks := make([]*domains.Task, 0)
	rows, err := t.db.QueryContext(ctx, fmt.Sprintf(query, conditions), params...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domains.ErrorNotFound
		}
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var (
			id          int64
			content     string
			userId      int64
			createdDate time.Time
		)
		rows.Scan(&id, &content, &userId, &createdDate)
		tasks = append(tasks, &domains.Task{
			Id:          id,
			Content:     content,
			UserId:      userId,
			CreatedDate: createdDate,
		})
	}

	return tasks, nil
}

func (t TaskRepositoryPostgresImpl) GetTaskById(ctx context.Context, request *domains.TaskByIdRequest) (*domains.Task, error) {
	var query = `
		SELECT "id", "content", "user_id", "created_date" 
		FROM 
			tasks
		WHERE 
			id = $1 AND user_id = $2
		LIMIT 1 FOR NO KEY UPDATE
	`

	var (
		id          int64
		content     string
		userId      int64
		createdDate time.Time
	)

	err := t.db.QueryRowContext(ctx, query, request.Id, request.UserId).Scan(&id, &content, &userId, &createdDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domains.ErrorNotFound
		}
		return nil, err
	}

	return &domains.Task{
		Id:          id,
		Content:     content,
		UserId:      userId,
		CreatedDate: createdDate,
	}, nil
}
