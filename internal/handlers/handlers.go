package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/manabie-com/togo/internal/driver"
	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/repository/postgres"
	"github.com/manabie-com/togo/internal/usecases/task"
	"github.com/manabie-com/togo/internal/usecases/user"
	"github.com/manabie-com/togo/internal/utils"
)

// Repository is the repository type for the handlers to use
type Repository struct {
	userUsecase user.UserUsecase
	taskUsecas  task.TaskUsecase
}

// Repo is the concrete repository used by the handlers
var Repo = &Repository{}

// SetRepoForHandlers set the repository for the handlers
func SetRepoForHandlers(r *Repository) {
	Repo = r
}

// NewRepo creates a new repository to be used by the handlers
func NewRepo(db *driver.DB) *Repository {
	dbrepo := postgres.NewPostgresRepository(db.SQL)

	return &Repository{
		userUsecase: user.NewUserUsecase(dbrepo),
		taskUsecas:  task.NewTaskUsecase(dbrepo),
	}
}

// Login is a handler for "/login" with method POST
func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	request := &models.User{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	username := utils.SqlNullString(request.Username)
	password := utils.SqlNullString(request.Password)

	if !m.userUsecase.ValidateUser(r.Context(), username, password) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid username/password",
		})
		return
	}

	user, err := m.userUsecase.GetUserByUserName(r.Context(), username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	token, err := m.userUsecase.GenerateToken(user.ID, user.MaxTaskPerDay)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	if err = json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
	}
}
