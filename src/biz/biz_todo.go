package biz

import (
	"errors"
	"sync"

	"github.com/HoangMV/todo/lib/log"
	"github.com/HoangMV/todo/src/models/entity"
	"github.com/HoangMV/todo/src/models/request"
)

func (biz *Biz) CreateTodo(req *request.CreateTodoReq) error {

	var (
		err1 error
		err2 error

		max, cur int
	)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		max, err1 = biz.dao.GetMaxUserTodoOneDay(req.UserID)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cur, err2 = biz.dao.CountUserTodoInCurrentDay(req.UserID)
	}()

	wg.Wait()

	if err1 != nil {
		return errors.New("retrieve max todo failed")
	}
	if err2 != nil {
		return errors.New("retrieve current todo failed")
	}

	if cur >= max {
		return errors.New("your todo count has reached its maximum")
	}

	todo := &entity.Todo{
		UserID:  req.UserID,
		Content: req.Content,
		Status:  entity.KStatusUncheck,
	}

	// Insert
	if err := biz.dao.InsertTodo(todo); err != nil {
		return errors.New("insert todo failed")
	}

	return nil
}

func (biz *Biz) UpdateTodo(req *request.UpdateTodoReq) error {
	todo := &entity.Todo{
		ID:      req.ID,
		Content: req.Content,
		Status:  entity.Status(req.Status),
	}

	if err := biz.dao.UpdateTodo(todo); err != nil {
		log.Get().Errorf("Biz::UpdateTodo UpdateTodo err:%v, data:%+v", err, todo)
		return errors.New("update todo failed")
	}

	return nil
}

func (biz *Biz) GetListUserTodo(req *request.GetTodosReq) ([]entity.Todo, error) {
	todos, err := biz.dao.SelectTodosByUserID(req.UserID, req.Size, req.Index)
	if err != nil {
		log.Get().Errorf("Biz::UpdateTodo SelectTodosByUserID err:%v, req:%+v", err, req)
		return nil, errors.New("get list todo failed")
	}

	return todos, nil
}
