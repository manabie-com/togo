package sqlstore

import (
	"context"
	"database/sql"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
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
