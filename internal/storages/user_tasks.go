package storages

import (
	"context"
	"time"

	"github.com/looplab/eventhorizon"

	user_tasks "github.com/manabie-com/togo/internal/services/user-tasks"

	"github.com/go-pg/pg"

	"github.com/google/uuid"
)

type userTask struct {
	tableName struct{} `sql:"user_tasks" pg:",discard_unknown_columns"`

	ID         uuid.UUID `sql:"id,type:uuid,pk"`
	Version    int       `sql:"version,type:int"`
	NumOfTasks int       `sql:"num_of_tasks"`
	CreatedAt  time.Time `sql:"created_at,type:timestamp with time zone"`
	UpdatedAt  time.Time `sql:"updated_at,type:timestamp with time zone"`
}

func convert(u userTask) *user_tasks.UserTask {
	return &user_tasks.UserTask{
		ID:         u.ID,
		Version:    u.Version,
		NumOfTasks: u.NumOfTasks,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}

type userTaskRepo struct {
	db *pg.DB
}

func NewUserTaskRepo(db *pg.DB) *userTaskRepo {
	return &userTaskRepo{db: db}
}

// Parent returns the parent read repository, if there is one.
// Useful for iterating a wrapped set of repositories to get a specific one.
func (r *userTaskRepo) Parent() eventhorizon.ReadRepo {
	return nil
}

// FindAll returns all entities in the repository.
func (r *userTaskRepo) FindAll(context.Context) ([]eventhorizon.Entity, error) {
	return []eventhorizon.Entity{}, nil
}

// Remove removes a entity by ID from the storage.
func (r *userTaskRepo) Remove(context.Context, uuid.UUID) error {
	return nil
}

// Save saves a entity in the storage.
func (r *userTaskRepo) Save(ctx context.Context, entity eventhorizon.Entity) error {
	if entity.EntityID() == uuid.Nil {
		return eventhorizon.RepoError{
			Err:       eventhorizon.ErrCouldNotSaveEntity,
			BaseErr:   eventhorizon.ErrMissingEntityID,
			Namespace: eventhorizon.NamespaceFromContext(ctx),
		}
	}

	m := entity.(*user_tasks.UserTask)

	mm := &userTask{
		ID:         m.ID,
		Version:    m.Version,
		NumOfTasks: m.NumOfTasks,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}

	db := r.db.WithContext(ctx)
	err := db.Update(mm)
	if err == pg.ErrNoRows {
		err = db.Insert(mm)
	}
	if err != nil {
		return eventhorizon.RepoError{
			Err:       eventhorizon.ErrCouldNotSaveEntity,
			BaseErr:   err,
			Namespace: eventhorizon.NamespaceFromContext(ctx),
		}
	}
	return nil
}

// Find returns an entity for an ID.
func (r *userTaskRepo) Find(ctx context.Context, id uuid.UUID) (eventhorizon.Entity, error) {
	m := userTask{
		ID: id,
	}
	err := r.db.WithContext(ctx).Select(&m)
	if err != nil {
		return nil, eventhorizon.RepoError{
			Err:       eventhorizon.ErrEntityNotFound,
			BaseErr:   err,
			Namespace: eventhorizon.NamespaceFromContext(ctx),
		}
	}

	return convert(m), nil
}
