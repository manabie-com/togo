package services

import (
	"errors"
	"time"
	"togo/models"
)

type TaskReq struct {
	ID          uint      `json:"id"`
	Content     string    `json:"content"`
	UserID      uint      `json:"user_id"`
	CreatedTask time.Time `json:"created_task"`

	PageNum  int
	PageSize int
}

type TaskRes struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
}

func toTaskRes(task *models.Task) *TaskRes {
	var taskRes = &TaskRes{
		ID:      task.ID,
		Content: task.Content,
	}
	return taskRes
}

func (obj *TaskReq) GetAll() ([]*TaskRes, error) {
	model := models.Task{}
	list, err := model.GetAll(obj.PageNum, obj.PageSize, getMaps())
	if err != nil {
		return nil, err
	}
	var res []*TaskRes
	for _, item := range list {
		res = append(res, toTaskRes(item))
	}
	return res, nil
}

func (obj *TaskReq) GetTotal() (int, error) {
	model := models.Task{}
	return model.GetTotal(getMaps())
}

func (obj *TaskReq) Add() (bool, error) {
	now := time.Now().Format("2006-01-02")
	//check total task of user
	totalTaskOfUser, err := (&models.Task{}).GetCountTaskByUser(obj.UserID, now)
	if err != nil {
		return false, err
	}
	if totalTaskOfUser > 10 {
		return false, errors.New("Can not create more task for day")
	}
	model := models.Task{
		Content:     obj.Content,
		UserID:      obj.UserID,
		CreatedTask: now,
	}
	_, err = model.Add()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (obj *TaskReq) Update() (bool, error) {
	model := models.Task{
		Content: obj.Content,
	}
	_, err := model.Update(obj.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (obj *TaskReq) Delete() (bool, error) {
	model := models.Task{
		Content: obj.Content,
	}
	_, err := model.Delete(obj.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}
