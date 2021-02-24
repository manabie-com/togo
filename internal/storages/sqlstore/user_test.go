package sqlstore

import (
	"database/sql"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"context"
)

func TestValidateUser(t *testing.T) {
	Convey("Add task", t, func() {
		store := setup()

		validate := store.ValidateUser(context.Background(),
			sql.NullString{String: "00001", Valid:  true},
			sql.NullString{String: "example", Valid:  true})
		So(validate, ShouldBeTrue)

		teardown(store)
	})
}
