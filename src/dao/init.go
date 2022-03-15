package dao

import (
	"time"

	"github.com/HoangMV/todo/lib/pgsql"
	"github.com/HoangMV/todo/src/models/entity"
	"github.com/jmoiron/sqlx"
	"github.com/patrickmn/go-cache"
)

type DAO struct {
	db    *sqlx.DB
	cache *cache.Cache
}

type IDAO interface {
	InsertUser(obj *entity.User) error
	InsertTodo(obj *entity.Todo) error
	InsertUserMaxTodo(obj *entity.UserTodoConfig) error
	UpdateTodo(obj *entity.Todo) error
	GetUserByUsername(username string) (*entity.User, error)
	SelectTodosByUserID(userID int, size, index int) ([]entity.Todo, error)
	CountUserTodoInCurrentDay(userID int) (int, error)
	GetMaxUserTodoOneDay(userID int) (int, error)

	GetTokenInCache(username string) int
	SetTokenToCache(token string, userID int)
}

func New() IDAO {
	var r IDAO = &DAO{
		db:    pgsql.Get(),
		cache: cache.New(1*time.Hour, 1*time.Hour+5*time.Minute),
	}
	return r
}
