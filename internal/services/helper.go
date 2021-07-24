package services

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/robfig/cron/v3"
)

type userAuthKey int8

func userIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(userAuthKey(0))
	id, ok := v.(string)
	return id, ok
}
func value(req *http.Request, p string) sql.NullString {
	return sql.NullString{
		String: req.FormValue(p),
		Valid:  true,
	}
}

var endOfToday time.Time
var todayStr string
var (
	defaultLocation = time.Local
)

func init() {
	calculateEndOfToday()
	c := cron.New(cron.WithLocation(defaultLocation))
	c.AddFunc("0 0 * * ?", func() {
		calculateEndOfToday()
	})
	c.Start()
}

func calculateEndOfToday() {
	yy, MM, dd := time.Now().In(defaultLocation).Date()
	today := time.Date(yy, MM, dd, 0, 0, 0, 0, defaultLocation)
	todayStr = today.Format("2006-01-02")
	endOfToday = today.AddDate(0, 0, 1)
	if endOfToday.Before(time.Now().In(defaultLocation)) {
		calculateEndOfToday()
	}
}
