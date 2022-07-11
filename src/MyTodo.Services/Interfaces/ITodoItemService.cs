using MyTodo.Services.ViewModels;
using MyTodo.Services.ViewModels.Common;
using MyTodo.Services.ViewModels.TodoItem;
using System;
using System.Collections.Generic;
using System.Text;

namespace MyTodo.Services.Interfaces
{
    public interface ITodoItemService
    {
        List<TodoItemViewModel> GetAll();

        List<TodoItemViewModel> GetAll(string keyword);

        PagedResult<TodoItemViewModel> GetAllPaging(TodoItemPagingRequest request);

        TodoItemViewModel GetById(int id);

        int Add(TodoItemViewModel request);

        int Update(TodoItemUpdateRequest request);

        int Remove(int id);

    }
}
