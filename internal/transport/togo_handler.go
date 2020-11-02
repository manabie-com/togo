package transport

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	togo "github.com/manabie-com/togo/internal"
	"github.com/manabie-com/togo/internal/entities"
	"github.com/manabie-com/togo/logging"
)

var logger = logging.Logger.With("package", "transport")

//TogoHandler represent the httphandler for togo
type TogoHandler struct {
	TogoUsecase togo.Usecase
	JWTKey      string
}

//NewTogoHandler constructor
func NewTogoHandler(mux *chi.Mux, us togo.Usecase, JWTKey string) {
	handler := &TogoHandler{
		TogoUsecase: us,
		JWTKey:      JWTKey,
	}
	// StripSlashes remove redundant slash in endpoint, example /login/ -> /login
	mux.Use(middleware.StripSlashes)
	// Add CORS
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
	}))

	mux.Post("/login", handler.GetAuthToken)
	mux.Group(func(r chi.Router) {
		r.Use(handler.authMiddleware)
		r.Get("/tasks", handler.ListTasks)
		r.Post("/tasks", handler.AddTask)

	})
}
func (t *TogoHandler) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ok bool
		r, ok = t.ValidToken(r)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// GetAuthToken handle login, check infomation if valid -> return a token for user
func (t *TogoHandler) GetAuthToken(resp http.ResponseWriter, req *http.Request) {
	user := entities.User{}
	decode := json.NewDecoder(req.Body)
	//ignore object keys which do not match any non-ignored, exported fields (in struct)
	decode.DisallowUnknownFields()

	err := decode.Decode(&user)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !t.TogoUsecase.ValidateUser(req.Context(), convertNullString(user.ID), convertNullString(user.Password)) {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "incorrect user_id/pwd",
		})
		return
	}
	resp.Header().Set("Content-Type", "application/json")

	token, err := t.CreateToken(user.ID)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]string{
		"data": token,
	})
}

//ListTasks get all task with a input as a date accodinate with UserID from Authorization header
func (t *TogoHandler) ListTasks(resp http.ResponseWriter, req *http.Request) {
	id, _ := userIDFromCtx(req.Context())
	createdDate := value(req, "created_date")
	if len(createdDate.String) != 10 {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(map[string]string{
			"message": "Invalid created_date",
		})
		return
	}
	tasks, err := t.TogoUsecase.RetrieveTasks(
		req.Context(),
		sql.NullString{
			String: id,
			Valid:  true,
		},
		createdDate,
	)

	resp.Header().Set("Content-Type", "application/json")

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string][]entities.Task{
		"data": tasks,
	})
}

//AddTask add task to the db accodinate content and Authorization header
func (t *TogoHandler) AddTask(resp http.ResponseWriter, req *http.Request) {
	task := entities.Task{}
	err := json.NewDecoder(req.Body).Decode(&task)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	validator := validator.New()
	err = validator.Struct(task)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	now := time.Now()
	userID, _ := userIDFromCtx(req.Context())
	maxTaskTodo, err := t.TogoUsecase.GetMaxTaskTodo(req.Context(), userID)
	if err != nil {
		logger.Errorw("Can't get max task todo", "detail", err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	tasks, err := t.TogoUsecase.RetrieveTasks(
		req.Context(),
		sql.NullString{
			String: userID,
			Valid:  true,
		},
		sql.NullString{String: now.Format("2006-01-02"), Valid: true},
	)
	if len(tasks) >= maxTaskTodo {
		resp.Header().Set("Content-Type", "application/json")
		resp.WriteHeader(http.StatusForbidden)
		json.NewEncoder(resp).Encode(map[string]string{
			"message": "You have exceeded the number of creating task per day.",
		})
		return
	}

	task.ID = uuid.New().String()
	task.UserID = userID
	task.CreatedDate = now.Format("2006-01-02")
	resp.Header().Set("Content-Type", "application/json")

	err = t.TogoUsecase.AddTask(req.Context(), task)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]entities.Task{
		"data": task,
	})
}

func value(req *http.Request, p string) sql.NullString {
	return sql.NullString{
		String: req.FormValue(p),
		Valid:  true,
	}
}
func convertNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: s, Valid: true}
}

//CreateToken create a token jwt with id
func (t *TogoHandler) CreateToken(id string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(t.JWTKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

//ValidToken check token is valid or not
func (t *TogoHandler) ValidToken(req *http.Request) (*http.Request, bool) {
	token := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(t.JWTKey), nil
	})
	if err != nil {
		logger.Error(err)
		return req, false
	}

	if !parsedToken.Valid {
		return req, false
	}

	id, ok := claims["user_id"].(string)
	if !ok {
		return req, false
	}

	req = req.WithContext(context.WithValue(req.Context(), userAuthKey(0), id))
	return req, true
}

type userAuthKey int8

func userIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(userAuthKey(0))
	id, ok := v.(string)
	return id, ok
}
