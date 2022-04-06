package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/manabie-com/togo/internal/driver"
	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/repository/postgres"
	"github.com/manabie-com/togo/internal/usecases/task"
	"github.com/manabie-com/togo/internal/usecases/user"
	"github.com/manabie-com/togo/internal/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
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

func (m *Repository) RetrieveTasks(w http.ResponseWriter, r *http.Request) {
	ok := true
	r, ok = m.isValidToken(r)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	userInfo := userInfoFromCtx(r.Context())

	tasks, err := m.taskUsecas.RetrieveTasks(r.Context(), userInfo.id, getFormValue(r, "created_date"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	if err = json.NewEncoder(w).Encode(map[string][]*models.Task{
		"data": tasks,
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
}

func (m *Repository) AddTask(w http.ResponseWriter, r *http.Request) {
	ok := true
	r, ok = m.isValidToken(r)

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	t := &models.Task{}
	if err := json.NewDecoder(r.Body).Decode(t); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if t.Detail == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "task detail must not be empty",
		})
		return
	}

	userInfo := userInfoFromCtx(r.Context())

	t.ID = uuid.New().String()
	t.UserID = uint(userInfo.id)
	t.CreatedDate = time.Now().Format("2006-01-02")

	isMaximum, err := m.taskUsecas.IsMaxTasksPerDay(r.Context(), t.UserID, userInfo.maxTaskPerDay, utils.SqlNullString(t.CreatedDate))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	if isMaximum {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "the maximum number of tasks for this user today is reached",
		})
		return
	}

	err = m.taskUsecas.AddTask(r.Context(), t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	if err = json.NewEncoder(w).Encode(map[string]*models.Task{
		"data": t,
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
	}
}

type userTokenCtxKey string
type userInfo struct {
	id            uint
	maxTaskPerDay uint
}

func (m *Repository) isValidToken(r *http.Request) (*http.Request, bool) {
	token := r.Header.Get("Authorization")

	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil || !t.Valid {
		return r, false
	}

	// note: it must be float64
	id, ok := claims["user_id"].(float64)
	if !ok {
		return r, false
	}

	// same here
	max_task_per_day, ok := claims["max_task_per_day"].(float64)
	if !ok {
		return r, false
	}

	tokenContext := userInfo{
		id:            uint(id),
		maxTaskPerDay: uint(max_task_per_day),
	}

	r = r.WithContext(context.WithValue(r.Context(), userTokenCtxKey("token"), tokenContext))

	return r, true
}

func userInfoFromCtx(ctx context.Context) userInfo {
	return ctx.Value(userTokenCtxKey("token")).(userInfo)
}

func getFormValue(r *http.Request, key string) sql.NullString {
	return utils.SqlNullString(r.FormValue(key))
}
