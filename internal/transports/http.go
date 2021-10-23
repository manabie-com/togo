package transports

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	svc "github.com/manabie-com/togo/internal/services"
)

// ToDoTrans implement HTTP server
type ToDoTrans struct {
	JWTKey string
	TodoSvc svc.ToDoService  
}

func (s *ToDoTrans) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL.Path)
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "*")

	if req.Method == http.MethodOptions {
		resp.WriteHeader(http.StatusOK)
		return
	}

	switch req.URL.Path {
	case "/login":
		s.getAuthToken(resp, req)
		return
	case "/tasks":
		var ok bool
		req, ok = s.validToken(req)
		if !ok {
			resp.WriteHeader(http.StatusUnauthorized)
			return
		}

		switch req.Method {
		case http.MethodGet:
			s.listTasks(resp, req)
		case http.MethodPost:
			s.addTask(resp, req)
		}
		return
	}
}


func (s *ToDoTrans) addTask(resp http.ResponseWriter, req *http.Request) {
	const limitedTasksPerDay = 5
	t := svc.Task{}
	err := json.NewDecoder(req.Body).Decode(&t)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	resp.Header().Set("Content-Type", "application/json")

	userID, _ := userIDFromCtx(req.Context())
	count, err := s.TodoSvc.CountTodayTasks(req.Context(), userID) 
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	if count >= limitedTasksPerDay {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "You've exceeded the maximum tasks created request per day",
		})
		return
	}

	t, err = s.TodoSvc.AddTask(req.Context(), userID, t.Content)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]svc.Task{
		"data": t,
	})
}

func (s *ToDoTrans) listTasks(resp http.ResponseWriter, req *http.Request) {
	id, _ := userIDFromCtx(req.Context())
	tasks, err := s.TodoSvc.ListTasks(
		req.Context(),
		id,
		req.FormValue("created_date"),
	)

	resp.Header().Set("Content-Type", "application/json")

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string][]svc.Task{
		"data": tasks,
	})
}


func (s *ToDoTrans) getAuthToken(resp http.ResponseWriter, req *http.Request) {
	id := req.FormValue("user_id")
	if !s.TodoSvc.ValidateUser(req.Context(),id , req.FormValue("password")) {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "incorrect user_id/pwd",
		})
		return
	}
	resp.Header().Set("Content-Type", "application/json")

	token, err := s.createToken(id)
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

func (s *ToDoTrans) createToken(id string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(s.JWTKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

//TODO: should move to middleware
func (s *ToDoTrans) validToken(req *http.Request) (*http.Request, bool) {
	token := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(s.JWTKey), nil
	})
	if err != nil {
		log.Println(err)
		return req, false
	}

	if !t.Valid {
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