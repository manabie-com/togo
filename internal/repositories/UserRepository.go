package repositories

import (
	"context"
	"github.com/manabie-com/togo/internal/models"
)

// ValidateUser returns tasks if match userID AND password
func (l *LiteDB) ValidateUser(ctx context.Context, userID, pwd string) bool {
	user := &models.User{}
	result := l.DB.Table("users").WithContext(ctx).Select("id").Where("id = ? AND password = ?", userID, pwd).First(&user)

	if result.Error != nil {
		return false
	}

	return true
}

func (l *LiteDB) GetUser(ctx context.Context, userID string) (*models.User, error) {
	user := &models.User{}
	result := l.DB.WithContext(ctx).Table("users").Select("id, max_todo").First(&user, "id = ?", userID)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
