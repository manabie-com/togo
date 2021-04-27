package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

func Debug(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(bytes))
	}
}

type UserAuthKey int8

func UserIDFromCtx(ctx context.Context) (uint64, bool) {
	v := ctx.Value(UserAuthKey(0))
	id, ok := v.(uint64)
	return id, ok
}
