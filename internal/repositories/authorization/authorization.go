package authorization

import (
	"os"
	"strconv"
	"time"

	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/usecases/authorization"
	"github.com/manabie-com/togo/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type repository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) authorization.AuthRepository {
	return &repository{
		DB: db,
	}
}

// ValidateUser implements auth.AuthRepository
func (r *repository) ValidateUser(username string) (bool, error) {
	if username == "" {
		return false, errors.New("Invalid input")
	}

	user := &models.User{}
	r.DB.Where("username = ?", username).First(user)

	if user == nil || user.Username == "" {
		return false, nil
	}

	return true, nil
}

func (r *repository) GenerateToken(userID, maxTaskPerday string) (*string, error) {
	if userID == "" {
		return nil, errors.New("Input empty")
	}

	// Init a map claim for storing essential info
	claims := jwt.MapClaims{}

	timeout, err := strconv.Atoi(os.Getenv("JWT_TIMEOUT"))
	if err != nil {
		return nil, err
	}

	//Set value for token
	claims["user_id"] = userID
	claims["max_task_per_day"] = maxTaskPerday

	// Must use 'exp' key for storing timeout info
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(timeout)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return nil, err
	}

	return utils.String(tokenString), nil
}
