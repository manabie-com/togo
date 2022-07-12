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
        public Assignment()
        {

        }
        public Assignment(int todoItemId, Guid userId, Guid assignedUser, DateTime assignedDate)
        {
            this.TodoItemId = todoItemId;
            this.UserId = userId;
            this.AssignedUser = assignedUser;
            this.AssignedDate = assignedDate;
        }
        public int TodoItemId { get; set; }

        public Guid UserId { get; set; }

        [ForeignKey("TodoItemId")]
        public virtual TodoItem TodoItem { get; set; }

        [ForeignKey("UserId")]
        public virtual AppUser User { get; set; }

        public DateTime AssignedDate { get; set; }

        public Guid AssignedUser { get; set; }

    }
}
