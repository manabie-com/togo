using MyTodo.Services.ViewModels.Common;
using System;
using System.Collections.Generic;
using System.Text;

namespace MyTodo.Services.ViewModels.TodoItem
{
    public class TodoItemPagingRequest : PagingRequestBase
    {
        public string Keyword { get; set; }

    }
}
