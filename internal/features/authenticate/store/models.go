package store

import (
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/features/authenticate"
)

// dbUser represent the structure we need for moving data
// between the app and the database.
type dbUser struct {
	ID           uuid.UUID `db:"id"`
	Email        string    `db:"email"`
	PasswordHash []byte    `db:"password_hash"`
}

func toFeatureUser(dbUsr dbUser) authenticate.User {
	usr := authenticate.User{
		ID:           dbUsr.ID,
		Email:        dbUsr.Email,
		PasswordHash: dbUsr.PasswordHash,
	}

	return usr
}
