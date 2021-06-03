package tasks

import (
	"encoding/json"
	"net/http"

	"github.com/manabie-com/togo/internal/consts"
	"github.com/manabie-com/togo/internal/models"
	timUtils "github.com/manabie-com/togo/internal/utils/time"
)

type ICreateRequest interface {
	Bind(*http.Request) error
	Validate() error
	ToModel(string) *models.Task
}

type CreateRequest struct {
	Content     string `json:"content"`
	CreatedDate string `json:"created_date"`
}

// Bind request data
func (r *CreateRequest) Bind(req *http.Request) error {

	err := json.NewDecoder(req.Body).Decode(r)
	defer req.Body.Close()
	if err != nil {
		return consts.ErrInvalidRequest
	}

	return nil
}

// Validate request data
func (r CreateRequest) Validate() error {
	//TODO: validate request data
	return nil
}

// Convert request into model DTO
func (r CreateRequest) ToModel(userID string) *models.Task {
	return &models.Task{
		Content:    r.Content,
		UserID:     userID,
		CreateDate: timUtils.CurrentDate(),
	}
}
