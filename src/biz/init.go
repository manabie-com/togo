package biz

import (
	"github.com/HoangMV/todo/src/dao"
)

type Biz struct {
	dao dao.IDAO
}

func New() *Biz {
	return &Biz{dao: dao.New()}
}

func NewWithDao(obj dao.IDAO) *Biz {
	return &Biz{dao: obj}
}
