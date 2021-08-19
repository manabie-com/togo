package resources

import (
	"github.com/google/uuid"
	models "github.com/manabie-com/togo/internal/models"
)

type TaskResource struct {
	ID          uuid.UUID 	`json:"id,string"`
	Content     string 		`json:"content"`
	UserID      uuid.UUID 	`json:"user_id,string"`
	CreatedDate string 		`json:"created_date"`
}

func ToTaskResource(task models.Task) TaskResource {
	return TaskResource{
		ID: task.ID, 
		Content: task.Content,
		UserID: task.UserID,
		CreatedDate: task.CreatedDate,
	}
}

func TaskResources(tasks []models.Task) []TaskResource {
	tasksResource := make([]TaskResource, len(tasks))

	for i, item := range tasks {
		tasksResource[i] = ToTaskResource(item)
	}

	return tasksResource
}
