package auth

import (
	"context"

	"github.com/looplab/eventhorizon"
)

const (
	jwtTokenKey = "jwtToken"
	uidKey      = "uid"
)

func init() {
	// Register the UID context.
	eventhorizon.RegisterContextMarshaler(MarshalUID)
}

func MarshalUID(ctx context.Context, vals map[string]interface{}) {
	if uid, ok := ctx.Value(uidKey).(UID); ok {
		vals[uidKey] = uid
	}
}
