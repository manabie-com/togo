package services

import (
	taskrepo "github.com/manabie-com/togo/internal/repository"
)

// ToDoService implement HTTP server
type ToDoService struct {
	JWTKey string
	Repo   *taskrepo.TaskRepository
}

// func (s *ToDoService) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
// 	log.Println(req.Method, req.URL.Path)
// 	resp.Header().Set("Access-Control-Allow-Origin", "*")
// 	resp.Header().Set("Access-Control-Allow-Headers", "*")
// 	resp.Header().Set("Access-Control-Allow-Methods", "*")

// 	if req.Method == http.MethodOptions {
// 		resp.WriteHeader(http.StatusOK)
// 		return
// 	}

// 	switch req.URL.Path {
// 	case "/login":
// 		//	s.getAuthToken(resp, req)
// 		return
// 	case "/tasks":
// 		var ok bool
// 		req, ok = s.validToken(req)
// 		if !ok {
// 			resp.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}

// 		switch req.Method {
// 		case http.MethodGet:
// 			//s.listTasks(resp, req)
// 		case http.MethodPost:
// 			s.addTask(resp, req)
// 		}
// 		return
// 	}
// }

// func (s *ToDoService) getAuthToken(resp http.ResponseWriter, req *http.Request) {
// 	id := value(req, "user_id")
// 	if !s.Store.ValidateUser(req.Context(), id, value(req, "password")) {
// 		resp.WriteHeader(http.StatusUnauthorized)
// 		json.NewEncoder(resp).Encode(map[string]string{
// 			"error": "incorrect user_id/pwd",
// 		})
// 		return
// 	}
// 	resp.Header().Set("Content-Type", "application/json")

// 	token, err := s.createToken(id.String)
// 	if err != nil {
// 		resp.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(resp).Encode(map[string]string{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	json.NewEncoder(resp).Encode(map[string]string{
// 		"data": token,
// 	})
// }

// func (s *ToDoService) listTasks(resp http.ResponseWriter, req *http.Request) {
// 	id, _ := userIDFromCtx(req.Context())
// 	tasks, err := s.Store.RetrieveTasks(
// 		req.Context(),
// 		sql.NullString{
// 			String: id,
// 			Valid:  true,
// 		},
// 		value(req, "created_date"),
// 	)

// 	resp.Header().Set("Content-Type", "application/json")

// 	if err != nil {
// 		resp.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(resp).Encode(map[string]string{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	json.NewEncoder(resp).Encode(map[string][]*storages.Task{
// 		"data": tasks,
// 	})
// }

// func (s *ToDoService) addTask(resp http.ResponseWriter, req *http.Request) {
// 	t := &entity.Task{}
// 	err := json.NewDecoder(req.Body).Decode(t)
// 	defer req.Body.Close()
// 	if err != nil {
// 		resp.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	now := time.Now()
// 	userID, _ := userIDFromCtx(req.Context())
// 	t.ID = uuid.New().String()
// 	t.UserID = userID
// 	t.CreatedDate = now.Format("2006-01-02")

// 	resp.Header().Set("Content-Type", "application/json")

// 	res, err := s.Repo.Add(t)

// 	fmt.Print(res.RowsAffected())

// 	//err = s.Store.AddTask(req.Context(), t)
// 	if err != nil {
// 		resp.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(resp).Encode(map[string]string{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	json.NewEncoder(resp).Encode(map[string]*entity.Task{
// 		"data": t,
// 	})
// }

// func value(req *http.Request, p string) sql.NullString {
// 	return sql.NullString{
// 		String: req.FormValue(p),
// 		Valid:  true,
// 	}
// }

// func (s *ToDoService) createToken(id string) (string, error) {
// 	atClaims := jwt.MapClaims{}
// 	atClaims["user_id"] = id
// 	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
// 	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
// 	token, err := at.SignedString([]byte(s.JWTKey))
// 	if err != nil {
// 		return "", err
// 	}
// 	return token, nil
// }
