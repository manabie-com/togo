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
        public TodoItem()
        {

        }
        public TodoItem(string title, string desc, int priority, TodoItemStatus status)
        {
            this.Title = title;
            this.Description = desc;
            this.Priority = priority;
            this.Status = status;
        }
        [MaxLength(255)]
        public string Title { get; set; }
        [MaxLength(500)]
        public string Description { get; set; }
        public int Priority { get; set; }
        public TodoItemStatus Status { get; set; }
        public virtual ICollection<Assignment> Assignments { get; set; }

    }
}
