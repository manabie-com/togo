package repositories

import (
	"context"
	"github.com/manabie-com/togo/internal/models"
)

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *LiteDB) RetrieveTasks(ctx context.Context, userID, createdDate string) ([]*models.Task, error) {
	var tasks []*models.Task
	result := l.DB.WithContext(ctx).Table("tasks").Select("id", "content", "user_id", "created_date").Where("user_id = ? AND created_date = ? ", userID, createdDate).Find(&tasks)

	if result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}

// AddTask adds a new task to DB
func (l *LiteDB) AddTask(ctx context.Context, t *models.Task) error {
	task := &models.Task{ID: t.ID, Content: t.Content, UserID: t.UserID, CreatedDate: t.CreatedDate}
	result := l.DB.WithContext(ctx).Create(&task)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
