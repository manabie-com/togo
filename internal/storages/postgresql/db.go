package postgresql

import (
	"context"
	"fmt"
	"github.com/banhquocdanh/togo/internal/storages"
	"github.com/go-pg/pg/v10"
	"github.com/go-redsync/redsync/v4"
	"time"
)

type PostgreSql struct {
	db *pg.DB
	rs *redsync.Redsync
}

type Option func(sql *PostgreSql)

func NewPostgreSQL(db *pg.DB, opts ...Option) *PostgreSql {
	p := &PostgreSql{db: db}
	for _, opt := range opts {
		opt(p)
	}

	return p
}

func WithRedSync(rs *redsync.Redsync) Option {
	return func(s *PostgreSql) {
		s.rs = rs
	}
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (p *PostgreSql) RetrieveTasks(ctx context.Context, userID, createdDate string) ([]*storages.Task, error) {
	var tasks []storages.Task
	err := p.db.Model(&tasks).Where("user_id = ?", userID).Where("created_date = ?", createdDate).Select()

	if err != nil {
		return nil, err
	}

	result := make([]*storages.Task, 0, len(tasks))
	for _, t := range tasks {
		result = append(result, &t)
	}
	return result, err
}

// insertTask insert a new task to DB
func (p *PostgreSql) insertTask(ctx context.Context, t *storages.Task) error {
	_, err := p.db.Model(t).Insert()
	return err
}

func (p *PostgreSql) countTask(userID, createdDate string) (int, error) {
	return p.db.Model(&storages.Task{
		UserID:      userID,
		CreatedDate: createdDate,
	}).Count()

}

func (p *PostgreSql) isValidToAddTask(userID, createdDate string) (bool, error) {
	user, err := p.GetUserByID(userID)
	if err != nil {
		return false, err
	}
	count, err := p.countTask(userID, createdDate)
	if err != nil {
		return false, err
	}

	return user.MaxTodo > uint(count), nil
}

func generateMutexKey(userID string) string {
	return fmt.Sprintf("add_task:%s", userID)
}

// addTask adds a new task to DB, with depend on user.MaxTodo
func (p *PostgreSql) addTask(ctx context.Context, t *storages.Task) error {
	ok, err := p.isValidToAddTask(t.UserID, t.CreatedDate)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("max todo task")
	}

	return p.insertTask(ctx, t)
}

// addTaskWithLock adds a new task to DB, with lock on userID
func (p *PostgreSql) addTaskWithLock(ctx context.Context, t *storages.Task) error {
	ok, err := p.isValidToAddTask(t.UserID, t.CreatedDate)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("max todo task")
	}

	mutex := p.rs.NewMutex(generateMutexKey(t.UserID), redsync.WithExpiry(8*time.Minute))
	if err := mutex.Lock(); err != nil {
		return err
	}
	defer mutex.Unlock()

	err = p.addTask(ctx, t)
	return err
}

func (p *PostgreSql) AddTask(ctx context.Context, t *storages.Task) error {
	if p.rs != nil {
		return p.addTaskWithLock(ctx, t)
	}
	return p.addTask(ctx, t)
}

// ValidateUser returns tasks if match userID AND password
func (p *PostgreSql) ValidateUser(ctx context.Context, userID, pwd string) bool {
	user, err := p.GetUserByID(userID)
	if err != nil {
		return false
	}

	return user.Password == pwd
}

func (p *PostgreSql) GetUserByID(userID string) (*storages.User, error) {
	user := &storages.User{
		ID: userID,
	}
	err := p.db.Model(user).WherePK().Select()
	return user, err
}
