package services

import (
	"database/sql"
	"encoding/json"
	usermodel "github.com/manabie-com/togo/internal/storages/user/model"
	"github.com/manabie-com/togo/pkg/common/crypto"
	"github.com/manabie-com/togo/up"
	"net/http"
)

func (s *ToDoService) Register(resp http.ResponseWriter, req *http.Request) {
	u := &up.RegisterRequest{}
	err := json.NewDecoder(req.Body).Decode(u)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	if u.MaxTodo == 0 {
		u.MaxTodo = s.maxTodo
	}

	userID := sql.NullString{
		String: u.ID,
		Valid:  true,
	}
	user, err := s.userstore.FindByID(req.Context(), userID)
	if err != nil {
		if err != sql.ErrNoRows {
			resp.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	if user != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "user exists with the given id",
		})
		return
	}

	err = s.userstore.Create(req.Context(), &usermodel.User{
		ID:       u.ID,
		Password: u.Password,
		MaxTodo:  u.MaxTodo,
	})
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	registerResponse := &up.RegisterResponse{
		ID:      u.ID,
		MaxTodo: u.MaxTodo,
	}

	json.NewEncoder(resp).Encode(map[string]interface{}{
		"data": registerResponse,
	})
}

func (s *ToDoService) Login(resp http.ResponseWriter, req *http.Request) {
	u := &usermodel.User{}
	err := json.NewDecoder(req.Body).Decode(u)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	userID := sql.NullString{
		String: u.ID,
		Valid:  true,
	}

	user, err := s.userstore.FindByID(req.Context(), userID)
	if err != nil {
		if err == sql.ErrNoRows {
			resp.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(resp).Encode(map[string]string{
				"error": "incorrect user_id/pwd",
			})
			return
		}
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	if user == nil || !crypto.CheckPasswordHash(u.Password, user.Password) {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "incorrect user_id/pwd",
		})
		return
	}
	resp.Header().Set("Content-Type", "application/json")

	token, err := s.createToken(u.ID)
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
