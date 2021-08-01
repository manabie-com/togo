package postgresql

import "togo/internal/entity"

func (task Task) MapToEntity() entity.Task {
	return entity.Task{
		ID:          task.ID,
		Content:     task.Content,
		UserID:      task.UserID,
		CreatedDate: task.CreatedDate,
		IsDone:      task.IsDone,
	}
}
