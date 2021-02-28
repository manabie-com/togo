package sqlstore

import (
	"database/sql"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"context"
)

func TestStore_FindByID(t *testing.T) {
	Convey("TestStore_FindByID_MatchID", t, func() {
		store := setup()

		user, err := store.FindByID(context.Background(), sql.NullString{String: "00001", Valid:  true})
		So(err, ShouldNotBeNil)
		So(user.ID, ShouldEqual, "00001")

		teardown(store)
	})
}
