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
				if !verifyUser(ctx) {
					code = http.StatusUnauthorized
					err = errors.New("invalid user")
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

func verifyUser(ctx *context.Context) bool {
	// Todo: implement verify user authentication and role
	// This action just to follow diagram
	// Micro-service should have security layer
	return true
}