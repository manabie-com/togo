using MyTodo.Data.Enums;
using MyTodo.Infrastructure;
using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace MyTodo.Data.Entities
{
    [Table("TodoItems")]

    public class TodoItem : EntityBase<int>
    {
        [MaxLength(255)]
        public string Title { get; set; }
        [MaxLength(500)]
        public string Description { get; set; }
        public int Priority { get; set; }
        public TodoItemStatus Status { get; set; }
    }
}
