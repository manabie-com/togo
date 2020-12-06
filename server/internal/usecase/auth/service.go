package auth

import (
	"context"
	"github.com/HoangVyDuong/togo/internal/storages/user"
)

type authService struct {}

//NewService create new service
func NewService() Service {
	return &authService{}
}

func (s *authService) Auth(ctx context.Context, user user.User, password string) error {
	return nil
}

func (s *authService) CreateToken(ctx context.Context, user user.User) string {
	return ""
}

//func (s *ToDoService) getAuthToken(resp http.ResponseWriter, req *http.Request) {
//	id := value(req, "user_id")
//	if !s.Store.ValidateUser(req.Context(), id, value(req, "password")) {
//		resp.WriteHeader(http.StatusUnauthorized)
//		json.NewEncoder(resp).Encode(map[string]string{
//			"error": "incorrect user_id/pwd",
//		})
//		return
//	}
//	resp.Header().Set("Content-Type", "application/json")
//
//	token, err := s.createToken(id.String)
//	if err != nil {
//		resp.WriteHeader(http.StatusInternalServerError)
//		json.NewEncoder(resp).Encode(map[string]string{
//			"error": err.Error(),
//		})
//		return
//	}
//
//	json.NewEncoder(resp).Encode(map[string]string{
//		"data": token,
//	})
//}