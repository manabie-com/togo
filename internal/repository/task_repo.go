package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/manabie-com/togo/internal/core/domain"
	"github.com/manabie-com/togo/internal/core/port"
	"github.com/manabie-com/togo/internal/shared"
	"github.com/manabie-com/togo/pkg/database"
)

func NewTaskRepository() port.TaskRepository {
	return new(taskRepo)
}

type taskRepo struct {
}

func (p *taskRepo) InitTables(ctx context.Context, conn database.Connection) error {
	_, err := conn.ExecContext(ctx, SQL_CREATE_USER_TABLE)
	if err != nil {
		return err
	}
	_, err = conn.ExecContext(ctx, SQL_CREATE_TASK_TABLE)
	return err
}

func (p *taskRepo) RetrieveTasks(ctx context.Context, conn database.Connection, userId, createdDate string) ([]*domain.Task, error) {
	rows, err := conn.QueryContext(
		ctx,
		SQL_TASK_GET_TASKS,
		userId,
		createdDate,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []*domain.Task
	for rows.Next() {
		task := new(domain.Task)
		err := rows.Scan(
			&task.Id,
			&task.Content,
			&task.UserId,
			&task.CreatedDate,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (p *taskRepo) AddTask(ctx context.Context, conn database.Connection, task *domain.Task) error {
	_, err := conn.ExecContext(
		ctx,
		SQL_TASK_ADD_TASK,
		task.Id,
		task.Content,
		task.UserId,
		task.CreatedDate,
	)
	return err
}

func (p *taskRepo) CheckIfCanAddTask(ctx context.Context, conn database.Connection, userId, checkedDate string) error {
	maxToDo, err := p.getUserMaxTodo(ctx, conn, userId)
	if err != nil {
		return err
	}
	countCreatedTasks, err := p.countTaskCreatedAt(ctx, conn, userId, checkedDate)
	if err != nil {
		return err
	}

	if countCreatedTasks >= maxToDo {
		return &shared.LimitedError{
			LimitedNumber: maxToDo,
		}
	}
	return nil
}

func (p *taskRepo) Login(ctx context.Context, conn database.Connection, username, password string) (string, error) {
	row := conn.QueryRowContext(
		ctx,
		SQL_TASK_GET_USER_ID,
		username,
		password,
	)
	var userId string
	err := row.Scan(
		&userId,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return "", err
	}
	return userId, nil
}

func (p *taskRepo) countTaskCreatedAt(ctx context.Context, conn database.Connection, userId, createdDate string) (int32, error) {
	row := conn.QueryRowContext(
		ctx,
		SQL_TASK_COUNT_TASKS_CREATED_AT,
		userId,
		createdDate,
	)
	var count int32
	err := row.Scan(
		&count,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return 0, err
	}
	return count, nil
}

func (p *taskRepo) getUserMaxTodo(ctx context.Context, conn database.Connection, userId string) (int32, error) {
	row := conn.QueryRowContext(
		ctx,
		SQL_TASK_GET_USER_MAX_TODO,
		userId,
	)
	var maxToDo int32
	err := row.Scan(
		&maxToDo,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("user does not exists")
		}
		return 0, err
	}
	return maxToDo, nil
}
