package api

import (
	"github.com/dinhquockhanh/togo/internal/app/auth"
	"github.com/dinhquockhanh/togo/internal/app/limit"
	"github.com/dinhquockhanh/togo/internal/app/task"
	"github.com/dinhquockhanh/togo/internal/app/user"
	"github.com/dinhquockhanh/togo/internal/pkg/config"
	db "github.com/dinhquockhanh/togo/internal/pkg/sql"
	"github.com/dinhquockhanh/togo/internal/pkg/token"
)

type Handler struct {
	task *task.Handler
	user *user.Handler
	auth *auth.Handler
}

func NewHandler() (*Handler, error) {
	cnn, err := db.NewSqlConnection(&config.All.DB)
	if err != nil {
		return nil, err
	}
	tokenizer, err := token.NewJwt(config.All.Token.SecretKey)
	if err != nil {
		return nil, err
	}

	// User
	ur := user.NewPostgresRepository(cnn)
	us := user.NewService(ur)
	uh := user.NewHandler(us)
	// Auth
	as := auth.NewService(us)
	ah := auth.NewHandler(tokenizer, as)
	// Limit
	lr := limit.NewPostgresRepository(cnn)
	ls := limit.NewService(lr)
	// Task
	tr := task.NewPostgresRepository(cnn)
	ts := task.NewService(tr)
	th := task.NewHandler(ts, us, ls)

	return &Handler{
		task: th,
		user: uh,
		auth: ah,
	}, nil
}
