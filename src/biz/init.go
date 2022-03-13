package biz

import (
	"github.com/HoangMV/togo/src/dao"
)

type Biz struct {
	dao *dao.DAO
}

func New() *Biz {
	return &Biz{dao.New()}
}
