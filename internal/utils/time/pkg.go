package time

import (
	"time"

	"github.com/manabie-com/togo/internal/consts"
)

func CurrentDate() string {
	return time.Now().Format(consts.DefaultDateFormat)
}
