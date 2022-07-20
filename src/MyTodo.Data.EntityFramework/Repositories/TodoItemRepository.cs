using MyTodo.Data.Entities;
using MyTodo.Data.Interfaces.Repositories;
using System;
using System.Collections.Generic;
using System.Text;

namespace MyTodo.Data.EntityFramework.Repositories
{
    public class TodoItemRepository:EFRepository<TodoItem, int>, ITodoItemRepository
    {
        public TodoItemRepository(MyTodoDbContext context) : base(context){}
    }
}
