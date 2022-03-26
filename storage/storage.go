package storage

import (
	"context"

	"github.com/luongdn/togo/database"
	"github.com/luongdn/togo/models"
)

type UserStore interface {
	CreateUser(ctx context.Context, user *models.User) error
}

type TaskStore interface {
	ListTasks(ctx context.Context, user_id string) ([]models.Task, error)
	CreateTask(ctx context.Context, user_id string, task *models.Task) error
}

type RuleStore interface {
	GetRule(ctx context.Context, user_id string, action models.UserAction) (models.Rule, error)
	CreateRule(ctx context.Context, user_id string, rule *models.Rule) error
	UpdateRule(ctx context.Context, user_id string, rule *models.Rule) error
}

func GetUserStore() UserStore {
	return NewUserStore(database.SQLdb)
}

func GetTaskStore() TaskStore {
	return NewTaskStore(database.SQLdb)
}

func GetRuleStore() RuleStore {
	return NewRuleStore(database.SQLdb, database.Rdb)
}
