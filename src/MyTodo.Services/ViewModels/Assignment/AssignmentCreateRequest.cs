using System;
using System.Collections.Generic;
using System.Text;

namespace MyTodo.Services.ViewModels.Assignment
{
    public class AssignmentCreateRequest
    {
        public int TodoItemId { get; set; }

        public Guid UserId { get; set; }

        public DateTime AssignedDate { get; set; }

        public Guid AssignedUser { get; set; }
    }
}
