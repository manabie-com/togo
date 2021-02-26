package postgres

import (
	"context"
	"fmt"
	"github.com/manabie-com/togo/internal/storages"
	"testing"
)

func TestPostgres_ValidateUser(t *testing.T) {
	ctx := context.Background()

	db := new(DatabaseMock)
	db.On("ValidateUser", ctx, "", "").Return(&storages.PgUser{}, nil)

	usr, err := db.ValidateUser(ctx, "", "")
	db.AssertExpectations(t)

	fmt.Println(usr, err)
}
