package mapper

import (
	"github.com/vchitai/togo/internal/models"
	"github.com/vchitai/togo/pb"
)

func Proto2ModelToDoEntry(todo *pb.ToDoEntry) *models.ToDo {
	return &models.ToDo{
		Content: todo.GetContent(),
	}
}

func Proto2ModelToDoEntryList(todoList []*pb.ToDoEntry) []*models.ToDo {
	var res = make([]*models.ToDo, 0, len(todoList))
	for _, todo := range todoList {
		res = append(res, Proto2ModelToDoEntry(todo))
	}
	return res
}
