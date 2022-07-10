using MyTodo.Infrastructure;
using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace MyTodo.Data.Entities
{
    [Table("Assignments")]
    public class Assignment : EntityBase<int>
    {
        public int TodoItemId { get; set; }

        public int UserId { get; set; }

        [ForeignKey("TodoItemId")]
        public virtual TodoItem TodoItem { get; set; }

        [ForeignKey("UserId")]
        public virtual User User { get; set; }
        public DateTime AssignedDate { get; set; }

    }
}
