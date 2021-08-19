package transports

import (
	"database/sql"
	"encoding/json"
	"net/http"

	repository "github.com/manabie-com/togo/internal/repositories"
	"github.com/manabie-com/togo/internal/services"
)

// ToDoService implement HTTP server
type ToDoLoginController struct {
	service *services.ToDoLoginService
}

func NewToDoLoginController(db *repository.DB, jwtKey string) *ToDoLoginController {
	return &ToDoLoginController{
		service: services.NewToDoLoginService(db, jwtKey),
	}
}

func (c *ToDoLoginController) GetAuthToken(resp http.ResponseWriter, req *http.Request) {
	id := value(req, "user_id")
	if !c.service.ValidateUser(req.Context(), c.service.Store.DB, id, value(req, "password")) {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "incorrect user_id/pwd",
		})
		return
	}
	resp.Header().Set("Content-Type", "application/json")

	token, err := c.service.CreateToken(id.String, c.service.JWTKey)
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

func value(req *http.Request, p string) sql.NullString {
	return sql.NullString{
		String: req.FormValue(p),
		Valid:  true,
	}
}
