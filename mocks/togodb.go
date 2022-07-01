package mocks

import "github.com/lawtrann/togo"

type TodoDB struct {
	GetUserByNameResp  togo.User
	GetUserByNameErr   error
	IsExceedPerDayResp bool
	IsExceedPerDayErr  error
	AddTodoByUserErr   error
}

func (tb TodoDB) GetUserByName(userName string) (*togo.User, error) {
	return &tb.GetUserByNameResp, tb.GetUserByNameErr
}

func (tb TodoDB) IsExceedPerDay(u togo.User) (bool, error) {
	return tb.IsExceedPerDayResp, tb.IsExceedPerDayErr
}

func (tb TodoDB) AddTodoByUser(u *togo.User, t *togo.Todo, uFlag bool) error {
	return tb.AddTodoByUserErr
}
