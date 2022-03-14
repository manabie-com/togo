package biz

import (
	"time"

	"github.com/HoangMV/togo/src/dao"

	"github.com/patrickmn/go-cache"
)

type Biz struct {
	dao   *dao.DAO
	cache *cache.Cache
}

func New() *Biz {
	return &Biz{
		dao:   dao.New(),
		cache: cache.New(1*time.Hour, 1*time.Hour+5*time.Minute),
	}
}
