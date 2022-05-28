package helper

import (
	"errors"
	"github.com/beego/beego/v2/server/web/context"
	"log"
	"net/http"
	"sync"
)

func BeforeTask(ctx *context.Context) (accepted bool, code int, err error) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			select {
			case <-ctx.Request.Context().Done():
				log.Println(ctx.Request.Context().Err())
				err = ctx.Request.Context().Err()
				wg.Done()
				return
			default:
				if valid, token := verifyToken(ctx); !valid {
					code = http.StatusUnauthorized
					err = errors.New("token unauthorized")
				} else if !checkRole(token, ctx.Request.Method) {
					code = http.StatusForbidden
					err = errors.New("user not allowed")
				} else {
					accepted = true
				}
				wg.Done()
				return
			}
		}
	}()
	wg.Wait()
	return
}

func verifyToken(header *context.Context) (bool, interface{}) {
	// Todo: implement verify user authentication
	return true, nil
}

func checkRole(token interface{}, method string) bool {
	// Todo: implement check user permission
	return true
}