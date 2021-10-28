package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/manabie-com/togo/internal/core/domain"
	"github.com/manabie-com/togo/internal/core/port"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ToDoService implement HTTP server
type ToDoService struct {
	jwtService  port.JwtService
	taskService port.TaskService
}

func (p *httpHandler) responseSuccess(c *gin.Context, data interface{}) {
	if data == nil {
		c.JSON(http.StatusOK, map[string]interface{}{})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}

func (p *httpHandler) responseError(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, map[string]string{
		"error": err.Error(),
	})
}

func (p *httpHandler) getUserIdFromRequest(c *gin.Context) (string, error) {
	token := c.GetHeader("Authorization")
	return p.jwtService.ParseToken(token)
}

func (p *httpHandler) login(c *gin.Context) {
	var req reqLogin
	err := c.BindJSON(&req)
	if err != nil {
		p.responseError(c, http.StatusBadRequest, err)
		return
	}

	userId, err := p.taskService.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		p.responseError(c, http.StatusInternalServerError, err)
		return
	}
	if len(userId) == 0 {
		p.responseError(c, http.StatusUnauthorized, errors.New("incorrect username or password"))
		return
	}

	token, err := p.jwtService.CreateToken(userId)
	if err != nil {
		p.responseError(c, http.StatusInternalServerError, err)
		return
	}
	p.responseSuccess(c, token)
}

func (p *httpHandler) getTasks(c *gin.Context) {
	userId, err := p.getUserIdFromRequest(c)
	if err != nil {
		p.responseError(c, http.StatusBadRequest, err)
		return
	}
	tasks, err := p.taskService.RetrieveTasks(c.Request.Context(), userId, c.Query("created_date"))
	if err != nil {
		p.responseError(c, http.StatusInternalServerError, err)
		return
	}
	p.responseSuccess(c, tasks)
}

func (p *httpHandler) addTask(c *gin.Context) {
	var req reqAddTask
	err := c.BindJSON(&req)
	if err != nil {
		p.responseError(c, http.StatusBadRequest, err)
		return
	}

	userId, err := p.getUserIdFromRequest(c)
	if err != nil {
		p.responseError(c, http.StatusBadRequest, err)
		return
	}

	now := time.Now()
	task := &domain.Task{
		Id:          uuid.New().String(),
		Content:     req.Content,
		UserId:      userId,
		CreatedDate: now.Format("2006-01-02"),
	}
	err = p.taskService.AddTask(c.Request.Context(), task)
	if err != nil {
		p.responseError(c, http.StatusInternalServerError, err)
		return
	}
	p.responseSuccess(c, task)
}

// func value(req *http.Request, p string) sql.NullString {
// 	return sql.NullString{
// 		String: req.FormValue(p),
// 		Valid:  true,
// 	}
// }

// func (p *ToDoService) createToken(id string) (string, error) {
// 	atClaims := jwt.MapClaims{}
// 	atClaims["user_id"] = id
// 	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
// 	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
// 	token, err := at.SignedString([]byte(p.JWTKey))
// 	if err != nil {
// 		return "", err
// 	}
// 	return token, nil
// }

// func (p *ToDoService) validToken(req *http.Request) (*http.Request, bool) {
// 	token := req.Header.Get("Authorization")

// 	claims := make(jwt.MapClaims)
// 	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
// 		return []byte(p.JWTKey), nil
// 	})
// 	if err != nil {
// 		log.Println(err)
// 		return req, false
// 	}

// 	if !t.Valid {
// 		return req, false
// 	}

// 	id, ok := claims["user_id"].(string)
// 	if !ok {
// 		return req, false
// 	}

// 	req = req.WithContext(context.WithValue(req.Context(), userAuthKey(0), id))
// 	return req, true
// }

// type userAuthKey int8

// func userIDFromCtx(ctx context.Context) (string, bool) {
// 	v := ctx.Value(userAuthKey(0))
// 	id, ok := v.(string)
// 	return id, ok
// }
