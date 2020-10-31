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
	"github.com/google/uuid"
	togo "github.com/manabie-com/togo/internal"
	"github.com/manabie-com/togo/internal/entities"
	"github.com/manabie-com/togo/logging"
)

var logger = logging.Logger.With("package", "transport")

//TogoHandler represent the httphandler for togo
type TogoHandler struct {
	togoUsecase togo.Usecase
	JWTKey      string
}

func NewTogoHandler(mux *chi.Mux, us togo.Usecase, JWTKey string) {
	handler := &TogoHandler{
		togoUsecase: us,
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

	mux.Post("/login", handler.getAuthToken)
	mux.Group(func(r chi.Router) {
		r.Use(handler.authMiddleware)
		r.Get("/tasks", handler.listTasks)
		r.Post("/tasks", handler.addTask)

	})
}
func (t *TogoHandler) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ok bool
		r, ok = t.validToken(r)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (t *TogoHandler) getAuthToken(resp http.ResponseWriter, req *http.Request) {
	user := entities.User{}
	decode := json.NewDecoder(req.Body)
	//ignore object keys which do not match any non-ignored, exported fields (in struct)
	decode.DisallowUnknownFields()
	err := decode.Decode(&user)
	if !t.togoUsecase.ValidateUser(req.Context(), convertNullString(user.ID), convertNullString(user.Password)) {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "incorrect user_id/pwd",
		})
		return
	}
	resp.Header().Set("Content-Type", "application/json")

	token, err := t.createToken(user.ID)
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

func (t *TogoHandler) listTasks(resp http.ResponseWriter, req *http.Request) {
	id, _ := userIDFromCtx(req.Context())
	tasks, err := t.togoUsecase.RetrieveTasks(
		req.Context(),
		sql.NullString{
			String: id,
			Valid:  true,
		},
		value(req, "created_date"),
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

func (t *TogoHandler) addTask(resp http.ResponseWriter, req *http.Request) {
	task := entities.Task{}
	err := json.NewDecoder(req.Body).Decode(&task)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	now := time.Now()
	userID, _ := userIDFromCtx(req.Context())
	task.ID = uuid.New().String()
	task.UserID = userID
	task.CreatedDate = now.Format("2006-01-02")

	resp.Header().Set("Content-Type", "application/json")

	err = t.togoUsecase.AddTask(req.Context(), task)
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
func (t *TogoHandler) createToken(id string) (string, error) {
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

func (t *TogoHandler) validToken(req *http.Request) (*http.Request, bool) {
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
