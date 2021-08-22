package resources

import (
	"github.com/google/uuid"
	models "github.com/manabie-com/togo/internal/models"
)

type AuthSignupResource struct {
	UserID      uuid.UUID 	`json:"user_id,string"`
	UserName    string 		`json:"name"`
	UserEmail 	string 		`json:"email"`
	AccessToken string 		`json:"access_token"`
}

func ToAuthSignupResource(user models.User) AuthSignupResource {
	return AuthSignupResource{
		UserID: user.ID, 
		UserName: user.Name,
		UserEmail: user.Email,
		AccessToken: user.AccessToken,
	}
}