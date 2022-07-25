package model

import (
	"context"
	"github.com/uptrace/bun"
	"net/url"
	"time"
)

type Task struct {
	bun.BaseModel `bun:"table:tasks,alias:task"`
	ID            int       `bun:"id,pk,autoincrement"`
	Content       string    `bun:"content"`
	PublishedDate time.Time `bun:"published_date"`
	Status        int       `bun:"status"`
	CreatedBy     int       `bun:"created_by,pk"`
	Owner         *User     `bun:"rel:belongs-to,join:created_by=id"`
	CreatedAt     time.Time `bun:"created_at,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,default:current_timestamp"`
}

type ErrExceedDailyLimit struct{}

func (e ErrExceedDailyLimit) Error() string {
	return "user is exceeded daily limit of tasks"
}

type ErrInvalidUser struct{}

func (e ErrInvalidUser) Error() string {
	return "invalid user"
}

// FindTasks /**
func FindTasks(ctx context.Context, db bun.IDB, params url.Values) ([]Task, error) {
	var tasks = make([]Task, 0)
	user, ok := ctx.Value("user").(*User)
	if !ok {
		return nil, &ErrInvalidUser{}
	}
	query := db.NewSelect().Model(&tasks)

	for k := range params {
		v := params.Get(k)

		if k == "published_date" {
			query.Where("published_date=?", v)
			continue
		}
	}

	query.Where("created_by=?", user.ID)

	err := query.Scan(ctx)

	return tasks, err
}

// FindOneTaskByID /**
func FindOneTaskByID(ctx context.Context, id int, db bun.IDB) (Task, error) {
	user, ok := ctx.Value("user").(*User)
	if !ok {
		return Task{}, &ErrInvalidUser{}
	}

	task := Task{
		ID: id,
	}

	err := db.NewSelect().
		Model(&task).
		Where("id=?", id).
		Where("created_by=?", user.ID).
		Scan(ctx)
	return task, err
}

// Create /**
func (task *Task) Create(ctx context.Context, db bun.IDB) error {
	user, ok := ctx.Value("user").(*User)
	if !ok {
		return &ErrInvalidUser{}
	}

	if user.IsExceededDailyLimit(ctx, db) {
		return &ErrExceedDailyLimit{}
	}

	task.PublishedDate = time.Now().UTC()
	task.CreatedAt = time.Now().UTC()
	task.UpdatedAt = time.Now().UTC()
	task.CreatedBy = user.ID

	if _, err := db.NewInsert().Model(task).Exec(ctx); err != nil {
		return err
	}
	return nil

}

// Update /**
func (task *Task) Update(ctx context.Context, db bun.IDB) error {
	// force current time on updated_at column
	task.UpdatedAt = time.Now().UTC()
	if _, err := db.NewUpdate().
		Model(task).
		ExcludeColumn("created_at", "id", "created_by", "published_date").
		Where("id=?", task.ID).
		Returning("*").
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

// Delete /**
func (task *Task) Delete(ctx context.Context, db bun.IDB) error {

	if _, err := db.NewDelete().
		Model(task).
		Where("id=?", task.ID).
		Exec(ctx); err != nil {
		return err
	}
	return nil
}
