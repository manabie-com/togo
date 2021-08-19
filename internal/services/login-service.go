package services

import (
	"context"
	"database/sql"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/internal/models"
	repository "github.com/manabie-com/togo/internal/repositories"
)

// ToDoService implement HTTP server
type ToDoLoginService struct {
	JWTKey string
	Store  *repository.DB
}

func NewToDoLoginService(db *repository.DB, jwtKey string) *ToDoLoginService {
	return &ToDoLoginService{
		JWTKey: jwtKey,
		Store:  db,
	}
}

func (s *ToDoLoginService) CreateToken(id string, JWTKey string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(JWTKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

// ValidateUser returns tasks if match userID AND password
func (s *ToDoLoginService) ValidateUser(ctx context.Context, store *sql.DB, userID, pwd sql.NullString) bool {
	stmt := `SELECT id FROM users WHERE id = $1 AND password = $2`
	row := store.QueryRowContext(ctx, stmt, userID, pwd)
	u := &models.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return false
	}

	return true
}
