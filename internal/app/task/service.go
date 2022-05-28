package task

import (
	"context"
	"fmt"
)

type Repository interface {
	CreateTask(ctx context.Context, req *CreateTaskReq) (*Task, error)
	AssignTask(ctx context.Context, req *AssignTaskReq) (*Task, error)
	GetByID(ctx context.Context, req *GetTaskByIdReq) (*Task, error)
	ListTasks(ctx context.Context, req *ListTasksReq) ([]*Task, error)
	Delete(ctx context.Context, req *DeleteTaskByIdReq) error
	CountTasksOfUserToDay(ctx context.Context, username string) (int, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateTask(ctx context.Context, req *CreateTaskReq) (*Task, error) {
	task, err := s.repo.CreateTask(ctx, req)

	if err != nil {
		return nil, fmt.Errorf("service.CreateTask: %w", err)
	}

	return task, nil
}

func (s *service) AssignTask(ctx context.Context, req *AssignTaskReq) (*Task, error) {
	task, err := s.repo.AssignTask(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("service.AssignTask: %w", err)
	}

	return task, nil
}
func (s *service) CountTasksOfUserToDay(ctx context.Context, username string) (int, error) {
	count, err := s.repo.CountTasksOfUserToDay(ctx, username)
	if err != nil {
		return 0, fmt.Errorf("service.CountTasksOfUserToDay: %w", err)
	}

	return count, nil
}

func (s *service) ListTasks(ctx context.Context, req *ListTasksReq) ([]*Task, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetByID(ctx context.Context, req *GetTaskByIdReq) (*Task, error) {
	task, err := s.repo.GetByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("service.GetTaskByID: %w", err)
	}

	return task, nil
}

func (s *service) Delete(ctx context.Context, req *DeleteTaskByIdReq) error {
	err := s.repo.Delete(ctx, req)

	if err != nil {
		return fmt.Errorf("service.CreateTask: %w", err)
	}

	return nil
}
